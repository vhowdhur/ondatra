// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package portgraph searches for a graph that satisfies a set of constraints.
package portgraph

import (
	"fmt"
	"regexp"
	"sort"

	log "github.com/golang/glog"
	"github.com/pkg/errors"
)

var (
	reAny = Regex(regexp.MustCompile(".*"))
)

// AbstractGraph is a representation of abstract nodes, ports, and edges.
type AbstractGraph struct {
	Desc  string // Description for the AbstractGraph for logging.
	Nodes []*AbstractNode
	Edges []*AbstractEdge
}

// fetchPort2Port2EdgeMap() returns the mapping of a src AbstractPort to a dst AbstractPort to an AbstractEdge.
func (g *AbstractGraph) fetchPort2Port2EdgeMap() (map[*AbstractPort]map[*AbstractPort]*AbstractEdge, error) {
	ret := make(map[*AbstractPort]map[*AbstractPort]*AbstractEdge)
	for _, e := range g.Edges {
		if _, ok := ret[e.Src]; ok {
			return nil, fmt.Errorf("port %q is attached to more than one other port; can only be attached to one", e.Src.Desc)
		}
		if _, ok := ret[e.Dst]; ok {
			return nil, fmt.Errorf("port %q is attached to more than one other port; can only be attached to one", e.Dst.Desc)
		}
		ret[e.Src] = map[*AbstractPort]*AbstractEdge{e.Dst: e}
		ret[e.Dst] = map[*AbstractPort]*AbstractEdge{e.Src: &AbstractEdge{e.Dst, e.Src}}
	}
	return ret, nil
}

// ConcreteGraph is a representation of concrete nodes, ports, and edges.
type ConcreteGraph struct {
	Desc  string // Description for the AbstractGraph for logging.
	Nodes []*ConcreteNode
	Edges []*ConcreteEdge
}

// fetchPort2Port2EdgeMap() returns the mapping of a src ConcretePort to a dst ConcretePort to a ConcreteEdge.
func (g *ConcreteGraph) fetchPort2Port2EdgeMap() map[*ConcretePort]map[*ConcretePort]*ConcreteEdge {
	ret := make(map[*ConcretePort]map[*ConcretePort]*ConcreteEdge)
	for _, e := range g.Edges {
		if m, ok := ret[e.Src]; ok {
			m[e.Dst] = e
		} else {
			ret[e.Src] = map[*ConcretePort]*ConcreteEdge{e.Dst: e}
		}
		if m, ok := ret[e.Dst]; ok {
			m[e.Src] = &ConcreteEdge{e.Dst, e.Src}
		} else {
			ret[e.Dst] = map[*ConcretePort]*ConcreteEdge{e.Src: &ConcreteEdge{e.Dst, e.Src}}
		}
	}
	return ret
}

// AbstractNode represent a vertex on an AbstractGraph.
// The AbstractNode has AbstractPorts and AbstractEdges that originate and terminate at AbstractPorts.
type AbstractNode struct {
	Desc        string                    // Description for the AbstractNode for logging.
	Ports       []*AbstractPort           // A list AbstractPorts that may be connected to another AbstractPort.
	Constraints map[string]NodeConstraint // A map of key value constraints of the AbstractNode.
}

// ConcreteNode represent a vertex on a ConcreteGraph.
// The ConcreteNode has ConcretePorts and ConcreteEdges that originate and terminate at ConcretePorts.
type ConcreteNode struct {
	Desc  string            // Description for the ConcreteNode for logging.
	Ports []*ConcretePort   // A list ConcretePorts that may be connected to another ConcretePort.
	Attrs map[string]string // A map of key value attributes of the ConctreteNode.
}

func (n *ConcreteNode) containsPort(p *ConcretePort) bool {
	for _, p2 := range n.Ports {
		if p == p2 {
			return true
		}
	}
	return false
}

// AbstractEdge represents a link from an source AbstractPort to a destination AbstractPort.
type AbstractEdge struct {
	Src, Dst *AbstractPort
}

func (e *AbstractEdge) String() string {
	if e.Dst == nil {
		return fmt.Sprintf("%q -> nil", e.Src.Desc)
	}
	return fmt.Sprintf("%q -> %q", e.Src.Desc, e.Dst.Desc)
}

// ConcreteEdge represents a link from an source ConcretePort to a destination ConcretePort.
type ConcreteEdge struct {
	Src, Dst *ConcretePort
}

// AbstractPort is a is point on an AbstractNode where an AbstractEdge may be attached.
type AbstractPort struct {
	Desc        string                    // Description for the AbstractPort for logging.
	Constraints map[string]PortConstraint // A map of key value constraints of the AbstractPort.
}

// ConcretePort is a is point on a ConcreteNode where a ConcreteEdge may be attached.
type ConcretePort struct {
	Desc  string            // Description for the ConcretePort for logging.
	Attrs map[string]string // A map of key value attributes of the ConcretePort.
}

// Assignment contains the AbstractNode -> ConcreteNode and AbstractPort -> ConcretePort mappings.
type Assignment struct {
	Node2Node map[*AbstractNode]*ConcreteNode
	Port2Port map[*AbstractPort]*ConcretePort
}

type maxAssignment struct {
	assignment         *Assignment
	numNodes, numPorts int
}

type deferredNodeConstraint func(map[*AbstractNode]*ConcreteNode) bool
type deferredPortConstraint func(map[*AbstractPort]*ConcretePort) (bool, *AbstractPort, error)

type solver struct {
	abstractGraph     *AbstractGraph
	superGraph        *ConcreteGraph
	absPort2Node      map[*AbstractPort]*AbstractNode                   // Map Port to Node it is part of.
	absNode2Node      map[*AbstractNode]map[*AbstractNode]int           // Map Node to Node and how many links there are.
	conPort2Port2Edge map[*ConcretePort]map[*ConcretePort]*ConcreteEdge // Cache the linked concrete ports to edge.
	absPort2Port2Edge map[*AbstractPort]map[*AbstractPort]*AbstractEdge // Cache the linked abstract ports to edge.

	maxAssign *maxAssignment // The "best" Assignment for reporting if the solve failed.

	// Constraint mappings. Deferred constraint can only be checked after all abstract elements are assigned.
	nodeConstraints         map[*AbstractNode]map[string]LeafConstraint
	portConstraints         map[*AbstractPort]map[string]LeafConstraint
	deferredNodeConstraints map[*AbstractNode][]deferredNodeConstraint
	deferredPortConstraints map[*AbstractPort][]deferredPortConstraint
}

// Solve returns as assignment from superGraph that satisfies abstractGraph.
func Solve(abstractGraph *AbstractGraph, superGraph *ConcreteGraph) (*Assignment, error) {
	solveErr := &SolveErr{absGraphDesc: abstractGraph.Desc, conGraphDesc: superGraph.Desc}
	if len(abstractGraph.Nodes) > len(superGraph.Nodes) {
		return nil, solveErr

	}
	if len(abstractGraph.Edges) > len(superGraph.Edges) {
		return nil, solveErr
	}

	// Preprocess data for the solve.
	// Map the AbstractPort to AbstractNode and initialize maps for maxSolve.
	absPort2Node := make(map[*AbstractPort]*AbstractNode)
	node2Node := make(map[*AbstractNode]*ConcreteNode)
	port2Port := make(map[*AbstractPort]*ConcretePort)
	for _, n := range abstractGraph.Nodes {
		node2Node[n] = nil
		for _, p := range n.Ports {
			absPort2Node[p] = n
			port2Port[p] = nil
		}
	}
	// Map how many links there are between each AbstractNode to calculate matches.
	absNode2Node := make(map[*AbstractNode]map[*AbstractNode]int)
	for _, e := range abstractGraph.Edges {
		srcNode, dstNode := e.Src, e.Dst
		if _, ok := absNode2Node[absPort2Node[srcNode]]; !ok {
			absNode2Node[absPort2Node[srcNode]] = make(map[*AbstractNode]int)
		}
		absNode2Node[absPort2Node[srcNode]][absPort2Node[dstNode]]++
	}

	absPort2Port2Edge, err := abstractGraph.fetchPort2Port2EdgeMap()
	if err != nil {
		return nil, solveErr
	}

	s := &solver{
		abstractGraph:           abstractGraph,
		superGraph:              superGraph,
		absNode2Node:            absNode2Node,
		absPort2Node:            absPort2Node,
		absPort2Port2Edge:       absPort2Port2Edge,
		maxAssign:               &maxAssignment{&Assignment{Node2Node: node2Node, Port2Port: port2Port}, 0, 0},
		nodeConstraints:         make(map[*AbstractNode]map[string]LeafConstraint),
		deferredNodeConstraints: make(map[*AbstractNode][]deferredNodeConstraint),
		portConstraints:         make(map[*AbstractPort]map[string]LeafConstraint),
		deferredPortConstraints: make(map[*AbstractPort][]deferredPortConstraint),
	}

	a, ok := s.solve()
	if !ok {
		solveErr.maxAssign = s.maxAssign.assignment
		return nil, solveErr
	}
	return a, nil
}

// solve provides a mapping of abstract nodes and ports to concrete nodes and ports.
// 1. Find all concrete nodes that match the abstract node (check attributes and # of ports).
// 2. For each mapping, check the concrete nodes and satisfy the edges of the abstract graph.
// 3. Assign concrete ports.
// solve always returns the max assignment and true if the solve completed.
func (s *solver) solve() (*Assignment, bool) {
	abs2ConNodes := make(map[*AbstractNode][]*ConcreteNode)
	s.processConstraints()
	for _, n := range s.abstractGraph.Nodes {
		nodes := s.matchNodes(n, s.superGraph)
		if len(nodes) == 0 {
			return nil, false
		}
		for _, conNode := range nodes {
			abs2ConNodes[n] = append(abs2ConNodes[n], conNode)
		}
	}

	s.conPort2Port2Edge = s.superGraph.fetchPort2Port2EdgeMap()

	// Generate all AbstractNode -> ConcreteNode mappings.
	abs2ConNodeChan, stop := genCombos(abs2ConNodes)
	defer stop()
	// Iterate through each mapping.
	// For each mapping, evaluate deferred constraints and attempt to match edges and assign ports.
	for abs2ConNode := range abs2ConNodeChan {
		if !s.matchDeferredNodes(abs2ConNode) {
			continue
		}
		// For this mapping, check that the ConcreteNodes can satisfy the AbstractEdges.
		if !s.checkEdges(abs2ConNode) {
			continue
		}
		if len(abs2ConNode) > s.maxAssign.numNodes {
			s.maxAssign.assignment.Node2Node = abs2ConNode
			s.maxAssign.numNodes = len(abs2ConNode)
		}

		// Since the edges can be satisfied, try to assign matching ports.
		abs2ConPort := s.assignEdges(abs2ConNode)
		if abs2ConPort == nil {
			continue
		}

		return &Assignment{abs2ConNode, abs2ConPort}, true
	}
	return nil, false
}

// checkEdges validates that the concrete nodes can satisfy the abstract edges.
func (s *solver) checkEdges(abs2ConNode map[*AbstractNode]*ConcreteNode) bool {
	// Iterate through the abstract edges.
	for absSrcNode, absDstNodes := range s.absNode2Node {
		for absDstNode, need := range absDstNodes {
			conSrcNode, conDstNode := abs2ConNode[absSrcNode], abs2ConNode[absDstNode]
			if !s.checkEdge(conSrcNode, conDstNode, need) {
				log.V(1).Infof("node %q cannot be assigned to %q because there are not enough edges to node %q assigned to %q; need %d",
					absSrcNode.Desc, conSrcNode.Desc, absDstNode.Desc, conDstNode.Desc, need)
				return false
			}
		}
	}
	return true
}

func (s *solver) checkEdge(conSrcNode, conDstNode *ConcreteNode, edgesNeeded int) bool {
	var count int
	// Verify whether the concrete Edges satisfy the abstract Edges.
	for _, conSrcPort := range conSrcNode.Ports {
		for p := range s.conPort2Port2Edge[conSrcPort] {
			if conDstNode.containsPort(p) {
				count++
				if count >= edgesNeeded {
					return true
				}
			}
		}
	}
	return false
}

type portAssigner struct {
	// Pre-processed data for this port solve. These maps should not change.
	absPort2Edge            map[*AbstractPort]*AbstractEdge
	abs2ConEdges            map[*AbstractEdge][]*ConcreteEdge
	deferredPortConstraints map[*AbstractPort][]deferredPortConstraint
	orderedAbsEdges         []*AbstractEdge

	// Stored data between recursions.
	deferredUntilPortConstraints map[*AbstractPort][]deferredPortConstraint
	abs2ConPorts                 map[*AbstractPort]*ConcretePort // Ports that have been assigned.
	assignedConPorts             map[*ConcretePort]struct{}
	assignQueue                  *assignQueue

	// maxAbs2ConPorts stores the best solve (that is not complete) for error reporting.
	maxAbs2ConPorts map[*AbstractPort]*ConcretePort
}

// assignQueue is the queue of Edges to solve next and stores the previous queue state.
type assignQueue struct {
	prev  *assignQueue
	queue []*AbstractEdge // Assign ports as edges; this significantly reduces search space.
}

func (aq *assignQueue) newQueue() *assignQueue {
	newQueue := &assignQueue{
		prev:  aq,
		queue: aq.queue,
	}
	return newQueue
}

func (pa *portAssigner) nextEdge() *AbstractEdge {
	for _, e := range pa.orderedAbsEdges {
		if _, ok := pa.abs2ConPorts[e.Src]; !ok {
			return e
		}
	}
	return nil
}

func (pa *portAssigner) assign() bool {
	var absEdge *AbstractEdge
	// If there are enqueued edges, pop one off.
	if len(pa.assignQueue.queue) != 0 {
		absEdge = pa.assignQueue.queue[len(pa.assignQueue.queue)-1]
		pa.assignQueue.queue = pa.assignQueue.queue[:len(pa.assignQueue.queue)-1]
	} else {
		absEdge = pa.nextEdge()
		if absEdge == nil {
			// Done.
			return true
		}

	}

	// Assign the port(s) in the edge.
	absSrcPort, absDstPort := absEdge.Src, absEdge.Dst

	// The ports that were assigned/enqueued.
	var assignedAbsPorts []*AbstractPort
	var assignedConPorts []*ConcretePort
	var enqueuedPorts []*AbstractPort
	rewindState := func() {
		for _, p := range assignedAbsPorts {
			delete(pa.abs2ConPorts, p)
		}
		for _, p := range assignedConPorts {
			delete(pa.assignedConPorts, p)
		}
		for _, p := range enqueuedPorts {
			delete(pa.deferredUntilPortConstraints, p)
		}
	}
	for _, conEdge := range pa.abs2ConEdges[absEdge] {
		conSrcPort, conDstPort := conEdge.Src, conEdge.Dst
		if _, ok := pa.assignedConPorts[conSrcPort]; ok {
			continue
		} else if _, ok := pa.assignedConPorts[conDstPort]; ok {
			continue
		}

		// Reset state for each iteration.
		queue := pa.assignQueue.newQueue()
		tryPortAssignment := func(queue *assignQueue) bool {
			assignedAbsPorts = nil
			assignedConPorts = nil
			enqueuedPorts = nil

			// Try assigning the ports.
			assign := func(a *AbstractPort, c *ConcretePort) {
				pa.abs2ConPorts[a] = c
				pa.assignedConPorts[c] = struct{}{}
				assignedAbsPorts = append(assignedAbsPorts, a)
				assignedConPorts = append(assignedConPorts, c)
			}
			assign(absSrcPort, conSrcPort)
			if absDstPort != nil {
				assign(absDstPort, conDstPort)
			}

			// Check whether the assigned ports work with constraints.
			checkAssignment := func(absPort *AbstractPort) bool {
				// Start with constraints from previously assigned ports that depend on the port currently being assigned.
				if dcs, ok := pa.deferredUntilPortConstraints[absPort]; ok {
					for _, dc := range dcs {
						ok, _, err := dc(pa.abs2ConPorts)
						if err != nil {
							// This is an impossibility and should not happen.
							log.Fatalf("could not solve: %v", err)
						} else if !ok {
							// This ConcreteEdge does not work with previous assignments.
							return false
						}
					}
				}
				// Check deferred constraints.
				// If the constraint depends on another port that has not been assigned yet, defer the constraint
				// and add associated edge to the queue.
				if constraints, ok := pa.deferredPortConstraints[absPort]; ok {
					for _, dc := range constraints {
						if ok, p, err := dc(pa.abs2ConPorts); err != nil {
							pa.deferredUntilPortConstraints[p] = append(pa.deferredUntilPortConstraints[p], dc)
							queue.queue = append(queue.queue, pa.absPort2Edge[p])
							enqueuedPorts = append(enqueuedPorts, p)
						} else if !ok {
							return false
						}
					}
				}
				return true
			}
			if ok := checkAssignment(absSrcPort); !ok {
				return false
			}
			if absDstPort != nil {
				if ok := checkAssignment(absDstPort); !ok {
					return false
				}
			}
			return true
		}
		// Try to assign the ports.
		// Evaluate deferredPortConstraints. If this assignment does not work with other currently
		// assigned ports, rewind the state and try the next assignment.
		if !tryPortAssignment(pa.assignQueue) {
			rewindState()
			continue
		}
		// This edge assignment works; recurse for the next edge.
		pa.assignQueue = queue
		if pa.assign() {
			return true
		}
		pa.assignQueue = pa.assignQueue.prev
		// Store the best assignment to report.
		if len(pa.abs2ConPorts) > len(pa.maxAbs2ConPorts) {
			pa.maxAbs2ConPorts = make(map[*AbstractPort]*ConcretePort, len(pa.abs2ConPorts))
			for absPort, conPort := range pa.abs2ConPorts {
				pa.maxAbs2ConPorts[absPort] = conPort
			}
		}
		rewindState()
	}
	// None of these Edges work with the current state.
	return false
}

func (s *solver) assignEdges(abs2ConNode map[*AbstractNode]*ConcreteNode) map[*AbstractPort]*ConcretePort {
	if len(s.absPort2Node) == 0 {
		return map[*AbstractPort]*ConcretePort{}
	}
	// Generate all possible edges for this node assignment.
	abs2ConEdgeCombos := make(map[*AbstractEdge][]*ConcreteEdge)
	absPort2Edge := make(map[*AbstractPort]*AbstractEdge)
	for absSrcNode, conSrcNode := range abs2ConNode {
		for _, absSrcPort := range absSrcNode.Ports {
			// Check attributes match, then check if matched ports link correctly.
			matchedConPorts := s.matchPorts(absSrcPort, conSrcNode.Ports)
			if len(matchedConPorts) == 0 {
				// No possible assignments for this port.
				return nil
			}
			absDstPort2Edge, ok := s.absPort2Port2Edge[absSrcPort]
			if !ok {
				// The abstract port is not part of an edge.
				// NOTE: Unfortunately, this causes some hackiness where we consider unlinked Port A
				// a different port from Port A linked to Port B.
				var conEdges []*ConcreteEdge
				for _, p := range matchedConPorts {
					conEdges = append(conEdges, &ConcreteEdge{p, nil})
				}
				e := &AbstractEdge{absSrcPort, nil}
				abs2ConEdgeCombos[e] = conEdges
				absPort2Edge[absSrcPort] = e
				continue
			}

			// The AbstractPort is part of an edge, so check the ConcretePorts link to the corresponding ConcreteNode.
			var absDstPort *AbstractPort
			var absEdge *AbstractEdge
			for absDstPort, absEdge = range absDstPort2Edge {
				break
			}
			conDstNode := abs2ConNode[s.absPort2Node[absDstPort]]
			for _, conSrcPort := range matchedConPorts {
				conDstPorts2Edge := s.conPort2Port2Edge[conSrcPort]
				for p, conEdge := range conDstPorts2Edge {
					if conDstNode.containsPort(p) && s.matchPort(absDstPort, p) {
						abs2ConEdgeCombos[absEdge] = append(abs2ConEdgeCombos[absEdge], conEdge)
						absPort2Edge[absSrcPort] = absEdge
						absPort2Edge[absDstPort] = absEdge
					}
				}
			}
			if len(abs2ConEdgeCombos[absEdge]) == 0 {
				return nil
			}
		}
	}
	var orderedEdges []*AbstractEdge
	for e := range abs2ConEdgeCombos {
		orderedEdges = append(orderedEdges, e)
	}
	sort.Slice(orderedEdges, func(i, j int) bool {
		ei, ej := orderedEdges[i], orderedEdges[j]
		eiWeight, ejWeight := len(ei.Src.Constraints), len(ej.Src.Constraints)
		if ei.Dst != nil {
			eiWeight += len(ei.Dst.Constraints)
		}
		if ej.Dst != nil {
			ejWeight += len(ej.Dst.Constraints)
		}
		return eiWeight > ejWeight
	})

	aq := &assignQueue{
		prev:  &assignQueue{},
		queue: []*AbstractEdge{},
	}
	pa := &portAssigner{
		abs2ConEdges:                 abs2ConEdgeCombos,
		absPort2Edge:                 absPort2Edge,
		deferredPortConstraints:      s.deferredPortConstraints,
		orderedAbsEdges:              orderedEdges,
		deferredUntilPortConstraints: make(map[*AbstractPort][]deferredPortConstraint, len(s.deferredPortConstraints)),
		abs2ConPorts:                 make(map[*AbstractPort]*ConcretePort, len(s.absPort2Node)),
		assignedConPorts:             make(map[*ConcretePort]struct{}, len(s.absPort2Node)),
		assignQueue:                  aq,
		maxAbs2ConPorts:              make(map[*AbstractPort]*ConcretePort, len(s.absPort2Node)),
	}
	if ok := pa.assign(); !ok {
		if len(pa.maxAbs2ConPorts) > s.maxAssign.numPorts {
			s.maxAssign.assignment.Node2Node = abs2ConNode
			maxPorts := make(map[*AbstractPort]*ConcretePort, len(s.absPort2Node))
			for p := range s.absPort2Node {
				if cp, ok := pa.maxAbs2ConPorts[p]; ok {
					maxPorts[p] = cp
				} else {
					maxPorts[p] = nil
				}
			}
			s.maxAssign.assignment.Port2Port = maxPorts
			s.maxAssign.numPorts = len(pa.maxAbs2ConPorts)
		}
		return nil
	}
	return pa.abs2ConPorts
}

// Matching code.
// processConstraints iterates through the abstract graph and processes constraints.
func (s *solver) processConstraints() {
	for _, n := range s.abstractGraph.Nodes {
		n := n
		nc := make(map[string]LeafConstraint)
		var dnc []deferredNodeConstraint
		addNodeConstraint := func(k string, c NodeConstraint, n *AbstractNode) {
			switch v := c.(type) {
			case *sameAsNode:
				dnc = append(dnc, func(abs2ConNode map[*AbstractNode]*ConcreteNode) bool {
					return v.match(k, n, abs2ConNode)
				})
				nc[k] = reAny
			case *notSameAsNode:
				dnc = append(dnc, func(abs2ConNode map[*AbstractNode]*ConcreteNode) bool {
					return v.match(k, n, abs2ConNode)
				})
				nc[k] = reAny
			case LeafConstraint:
				nc[k] = v
			default:
				log.Fatalf("Unrecognized constraint type %T for node matching", v)
			}
		}
		for k, c := range n.Constraints {
			if v, ok := c.(*andNode); ok {
				for _, c := range v.nlcs {
					addNodeConstraint(k, c, n)
				}
			} else {
				addNodeConstraint(k, c, n)
			}
		}
		s.nodeConstraints[n] = nc
		s.deferredNodeConstraints[n] = dnc
		for _, p := range n.Ports {
			p := p
			pc := make(map[string]LeafConstraint)
			var dpc []deferredPortConstraint
			addPortConstraint := func(k string, c PortConstraint, p *AbstractPort) {
				switch v := c.(type) {
				case *sameAsPort:
					dpc = append(dpc, func(abs2ConPort map[*AbstractPort]*ConcretePort) (bool, *AbstractPort, error) {
						if _, ok := abs2ConPort[v.p]; !ok {
							return false, v.p, errors.Errorf("port %q not assigned", v.p.Desc)
						}
						return v.match(k, p, abs2ConPort), nil, nil
					})
					pc[k] = reAny
				case *notSameAsPort:
					dpc = append(dpc, func(abs2ConPort map[*AbstractPort]*ConcretePort) (bool, *AbstractPort, error) {
						if _, ok := abs2ConPort[v.p]; !ok {
							return false, v.p, errors.Errorf("port %q not assigned", v.p.Desc)
						}
						return v.match(k, p, abs2ConPort), nil, nil
					})
					pc[k] = reAny
				case LeafConstraint:
					pc[k] = v
				default:
					log.Fatalf("Unrecognized constraint type %T for port matching", v)
				}
			}
			for k, c := range p.Constraints {
				if v, ok := c.(*andPort); ok {
					for _, c := range v.plcs {
						addPortConstraint(k, c, p)
					}
				} else {
					addPortConstraint(k, c, p)
				}
			}
			s.portConstraints[p] = pc
			s.deferredPortConstraints[p] = dpc
		}
	}
}

// matchDeferredNodes checks whether the proposed node mapping fulfills constraints.
func (s *solver) matchDeferredNodes(abs2ConNode map[*AbstractNode]*ConcreteNode) bool {
	for _, dncs := range s.deferredNodeConstraints {
		for _, dnc := range dncs {
			if !dnc(abs2ConNode) {
				return false
			}
		}
	}
	return true
}

// matchNodes returns all ConcreteNodes the ConcreteGraph that match the AbstractNode.
func (s *solver) matchNodes(abs *AbstractNode, superGraph *ConcreteGraph) []*ConcreteNode {
	match := func(n *ConcreteNode) bool {
		// Check that there are at least enough Ports to satisfy potential Edges.
		if len(abs.Ports) > len(n.Ports) {
			return false
		}
		for k, c := range s.nodeConstraints[abs] {
			s, ok := n.Attrs[k]
			if !c.match(s, ok) {
				return false
			}
		}
		return true
	}

	var nodes []*ConcreteNode
	for _, n := range superGraph.Nodes {
		if match(n) {
			nodes = append(nodes, n)
		}
	}
	return nodes
}

// matchDeferredPort checks the deferred constraints against the port.
func (s *solver) matchDeferredPort(port *AbstractPort, abs2ConPort map[*AbstractPort]*ConcretePort, attr string, constraint PortConstraint) bool {
	switch v := constraint.(type) {
	case *sameAsPort:
		if !constraint.(*sameAsPort).match(attr, port, abs2ConPort) {
			return false
		}
	case *notSameAsPort:
		if !constraint.(*notSameAsPort).match(attr, port, abs2ConPort) {
			return false
		}
	default:
		log.Fatalf("Unrecognized constraint type %T for port matching", v)
	}
	return true

}

func (s *solver) matchPort(abs *AbstractPort, p *ConcretePort) bool {
	for k, c := range s.portConstraints[abs] {
		s, ok := p.Attrs[k]
		if !c.match(s, ok) {
			return false
		}
	}
	return true
}

// matchPorts returns all ConcretePorts that match the AbstractPort.
func (s *solver) matchPorts(abs *AbstractPort, con []*ConcretePort) []*ConcretePort {
	var ports []*ConcretePort
	for _, p := range con {
		if s.matchPort(abs, p) {
			ports = append(ports, p)
		}
	}
	return ports
}
