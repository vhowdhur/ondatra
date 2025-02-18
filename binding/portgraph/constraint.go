// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package portgraph

import (
	"regexp"

	log "github.com/golang/glog"
)

// NodeConstraint is a constraint on a Node.
type NodeConstraint interface {
	isNodeConstraint()
}

// LeafNodeConstraint is a leaf constraint on a Node.
type LeafNodeConstraint interface {
	isLeafConstraint()
	NodeConstraint
}

// PortConstraint is a constraint on a Port.
type PortConstraint interface {
	isPortConstraint()
}

// LeafPortConstraint is a leaf constraint on a Port.
type LeafPortConstraint interface {
	isLeafConstraint()
	PortConstraint
}

// LeafConstraint is a constraint checker that checks an input matches based on some matching criteria.
type LeafConstraint interface {
	universalLeafConstraint
	match(string, bool) bool
}

type universalLeafConstraint interface {
	LeafNodeConstraint
	LeafPortConstraint
}

type equal struct {
	universalLeafConstraint
	s string
}

func (c *equal) match(s string, ok bool) bool {
	return ok && s == c.s
}

type notEqual struct {
	universalLeafConstraint
	s string
}

func (c *notEqual) match(s string, ok bool) bool {
	return ok && s != c.s
}

type regex struct {
	universalLeafConstraint
	re *regexp.Regexp
}

func (c *regex) match(s string, ok bool) bool {
	return ok && c.re.MatchString(s)
}

type notRegex struct {
	universalLeafConstraint
	re *regexp.Regexp
}

func (c *notRegex) match(s string, ok bool) bool {
	return ok && !c.re.MatchString(s)
}

func attrsNodePair(k string, n1, n2 *AbstractNode, abs2ConNode map[*AbstractNode]*ConcreteNode) (string, string, bool) {
	cn1, ok := abs2ConNode[n1]
	if !ok {
		log.Fatalf("node %q does not have an assignment", n1.Desc)
	}
	v1, ok := cn1.Attrs[k]
	if !ok {
		return "", "", false
	}
	cn2, ok := abs2ConNode[n2]
	if !ok {
		log.Fatalf("node %q does not have an assignment", n2.Desc)
	}
	v2, ok := cn2.Attrs[k]
	if !ok {
		return "", "", false
	}
	return v1, v2, ok
}

func attrsPortPair(k string, p1, p2 *AbstractPort, abs2ConPort map[*AbstractPort]*ConcretePort) (string, string, bool) {
	cp1, ok := abs2ConPort[p1]
	if !ok {
		log.Fatalf("port %q does not have an assignement", p1.Desc)
	}
	v1, ok := cp1.Attrs[k]
	if !ok {
		return "", "", false
	}
	cp2, ok := abs2ConPort[p2]
	if !ok {
		log.Fatalf("port %q does not have an assignment", p2.Desc)
	}
	v2, ok := cp2.Attrs[k]
	if !ok {
		return "", "", false
	}
	return v1, v2, ok
}

type sameAsNode struct {
	LeafNodeConstraint
	n *AbstractNode
}

func (c *sameAsNode) match(k string, n *AbstractNode, abs2ConNode map[*AbstractNode]*ConcreteNode) bool {
	v1, v2, ok := attrsNodePair(k, n, c.n, abs2ConNode)
	return ok && v1 == v2
}

type sameAsPort struct {
	LeafPortConstraint
	p *AbstractPort
}

func (c *sameAsPort) match(k string, p *AbstractPort, abs2ConPort map[*AbstractPort]*ConcretePort) bool {
	v1, v2, ok := attrsPortPair(k, p, c.p, abs2ConPort)
	return ok && v1 == v2
}

type notSameAsNode struct {
	LeafNodeConstraint
	n *AbstractNode
}

func (c *notSameAsNode) match(k string, n *AbstractNode, abs2ConNode map[*AbstractNode]*ConcreteNode) bool {
	v1, v2, ok := attrsNodePair(k, n, c.n, abs2ConNode)
	return ok && v1 != v2
}

type notSameAsPort struct {
	LeafPortConstraint
	p *AbstractPort
}

func (c *notSameAsPort) match(k string, p *AbstractPort, abs2ConPort map[*AbstractPort]*ConcretePort) bool {
	v1, v2, ok := attrsPortPair(k, p, c.p, abs2ConPort)
	return ok && v1 != v2
}

type andNode struct {
	NodeConstraint
	nlcs []LeafNodeConstraint
}

type andPort struct {
	PortConstraint
	plcs []LeafPortConstraint
}

// AndNode returns a constraint that contains other node constraints.
// An and constraint is a leaf constraint that cannot contain other leaves.
func AndNode(nlcs ...LeafNodeConstraint) NodeConstraint {
	return &andNode{nlcs: nlcs}
}

// AndPort returns a constraint that contains other port constraints.
// An and constraint is a leaf constraint that cannot contain other leaves.
func AndPort(plcs ...LeafPortConstraint) PortConstraint {
	return &andPort{plcs: plcs}
}

// Equal returns a constraint that an attribute is equal to the specified string.
func Equal(s string) LeafConstraint {
	return &equal{s: s}
}

// NotEqual returns a constraint that an attribute is not equal to the specified string.
func NotEqual(s string) LeafConstraint {
	return &notEqual{s: s}
}

// Regex returns a constraint that an attribute matches the specified regex.
func Regex(re *regexp.Regexp) LeafConstraint {
	return &regex{re: re}
}

// NotRegex returns a constraint that an attribute does not match the specified regex.
func NotRegex(re *regexp.Regexp) LeafConstraint {
	return &notRegex{re: re}
}

// SameAsNode returns a constraint that an attribute has the same value as the specified Node.
func SameAsNode(n *AbstractNode) LeafNodeConstraint {
	return &sameAsNode{n: n}
}

// NotSameAsNode returns a constraint that an attribute has a different value from the specified Node.
func NotSameAsNode(n *AbstractNode) LeafNodeConstraint {
	return &notSameAsNode{n: n}
}

// SameAsPort returns a constraint that an attribute has the same value as the specified Port.
func SameAsPort(p *AbstractPort) LeafPortConstraint {
	return &sameAsPort{p: p}
}

// NotSameAsPort returns a constraint that an attribute has a different value from the specified Port.
func NotSameAsPort(p *AbstractPort) LeafPortConstraint {
	return &notSameAsPort{p: p}
}
