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

package knebind

import (
	"google.golang.org/grpc"
	"github.com/openconfig/ondatra/binding"

	bpb "github.com/openconfig/gnoi/bgp"
	cmpb "github.com/openconfig/gnoi/cert"
	dpb "github.com/openconfig/gnoi/diag"
	frpb "github.com/openconfig/gnoi/factory_reset"
	fpb "github.com/openconfig/gnoi/file"
	hpb "github.com/openconfig/gnoi/healthz"
	ipb "github.com/openconfig/gnoi/interface"
	lpb "github.com/openconfig/gnoi/layer2"
	mpb "github.com/openconfig/gnoi/mpls"
	ospb "github.com/openconfig/gnoi/os"
	otpb "github.com/openconfig/gnoi/otdr"
	spb "github.com/openconfig/gnoi/system"
	wrpb "github.com/openconfig/gnoi/wavelength_router"
)

type gnoiClients struct {
	*binding.AbstractGNOIClients
	conn *grpc.ClientConn
}

func (g *gnoiClients) BGP() bpb.BGPClient {
	return bpb.NewBGPClient(g.conn)
}

func (g *gnoiClients) CertificateManagement() cmpb.CertificateManagementClient {
	return cmpb.NewCertificateManagementClient(g.conn)
}

func (g *gnoiClients) Diag() dpb.DiagClient {
	return dpb.NewDiagClient(g.conn)
}

func (g *gnoiClients) FactoryReset() frpb.FactoryResetClient {
	return frpb.NewFactoryResetClient(g.conn)
}

func (g *gnoiClients) File() fpb.FileClient {
	return fpb.NewFileClient(g.conn)
}

func (g *gnoiClients) Healthz() hpb.HealthzClient {
	return hpb.NewHealthzClient(g.conn)
}

func (g *gnoiClients) Interface() ipb.InterfaceClient {
	return ipb.NewInterfaceClient(g.conn)
}

func (g *gnoiClients) Layer2() lpb.Layer2Client {
	return lpb.NewLayer2Client(g.conn)
}

func (g *gnoiClients) MPLS() mpb.MPLSClient {
	return mpb.NewMPLSClient(g.conn)
}

func (g *gnoiClients) OS() ospb.OSClient {
	return ospb.NewOSClient(g.conn)
}

func (g *gnoiClients) OTDR() otpb.OTDRClient {
	return otpb.NewOTDRClient(g.conn)
}

func (g *gnoiClients) System() spb.SystemClient {
	return spb.NewSystemClient(g.conn)
}

func (g *gnoiClients) WavelengthRouter() wrpb.WavelengthRouterClient {
	return wrpb.NewWavelengthRouterClient(g.conn)
}
