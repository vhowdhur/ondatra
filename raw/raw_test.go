// Copyright 2019 Google LLC
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

package raw

import (
	"golang.org/x/net/context"
	"errors"
	"strings"
	"testing"

	"google.golang.org/grpc"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/fakebind"
	"github.com/openconfig/testt"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
	grpb "github.com/openconfig/gribi/v1/proto/service"
	p4pb "github.com/p4lang/p4runtime/go/p4/v1"
)

var dut = &fakebind.DUT{AbstractDUT: &binding.AbstractDUT{&binding.Dims{Name: "fakeDUT"}}}

func TestGNMI(t *testing.T) {
	gnmi := NewDUTAPIs(dut).GNMI()

	t.Run("error", func(t *testing.T) {
		wantErr := "bad gnmi"
		dut.DialGNMIFn = func(context.Context, ...grpc.DialOption) (gpb.GNMIClient, error) {
			return nil, errors.New(wantErr)
		}
		gotErr := testt.ExpectFatal(t, func(t testing.TB) {
			gnmi.New(t)
		})
		if !strings.Contains(gotErr, wantErr) {
			t.Errorf("New(t) got err %v, want %v", gotErr, wantErr)
		}
		gotErr = testt.ExpectFatal(t, func(t testing.TB) {
			gnmi.Default(t)
		})
		if !strings.Contains(gotErr, wantErr) {
			t.Errorf("Default(t) got err %v, want %v", gotErr, wantErr)
		}
	})

	t.Run("success", func(t *testing.T) {
		want := struct{ gpb.GNMIClient }{}
		dut.DialGNMIFn = func(context.Context, ...grpc.DialOption) (gpb.GNMIClient, error) {
			return want, nil
		}
		if got := gnmi.New(t); got != want {
			t.Errorf("New(t) got %v, want %v", got, want)
		}
		if got := gnmi.Default(t); got != want {
			t.Errorf("Default(t) got %v, want %v", got, want)
		}
	})
}

func TestGNOI(t *testing.T) {
	gnoi := NewDUTAPIs(dut).GNOI()

	t.Run("error", func(t *testing.T) {
		wantErr := "bad gnoi"
		dut.DialGNOIFn = func(context.Context, ...grpc.DialOption) (binding.GNOIClients, error) {
			return nil, errors.New(wantErr)
		}
		gotErr := testt.ExpectFatal(t, func(t testing.TB) {
			gnoi.New(t)
		})
		if !strings.Contains(gotErr, wantErr) {
			t.Errorf("New(t) got err %v, want %v", gotErr, wantErr)
		}
		gotErr = testt.ExpectFatal(t, func(t testing.TB) {
			gnoi.Default(t)
		})
		if !strings.Contains(gotErr, wantErr) {
			t.Errorf("Default(t) got err %v, want %v", gotErr, wantErr)
		}
	})

	t.Run("success", func(t *testing.T) {
		want := struct{ binding.GNOIClients }{}
		dut.DialGNOIFn = func(context.Context, ...grpc.DialOption) (binding.GNOIClients, error) {
			return want, nil
		}
		if got := gnoi.New(t); got != want {
			t.Errorf("New(t) got %v, want %v", got, want)
		}
		if got := gnoi.Default(t); got != want {
			t.Errorf("Default(t) got %v, want %v", got, want)
		}
	})
}

func TestGRIBI(t *testing.T) {
	gribi := NewDUTAPIs(dut).GRIBI()

	t.Run("error", func(t *testing.T) {
		wantErr := "bad gribi"
		dut.DialGRIBIFn = func(context.Context, ...grpc.DialOption) (grpb.GRIBIClient, error) {
			return nil, errors.New(wantErr)
		}
		gotErr := testt.ExpectFatal(t, func(t testing.TB) {
			gribi.New(t)
		})
		if !strings.Contains(gotErr, wantErr) {
			t.Errorf("New(t) got err %v, want %v", gotErr, wantErr)
		}
		gotErr = testt.ExpectFatal(t, func(t testing.TB) {
			gribi.Default(t)
		})
		if !strings.Contains(gotErr, wantErr) {
			t.Errorf("Default(t) got err %v, want %v", gotErr, wantErr)
		}
	})

	t.Run("success", func(t *testing.T) {
		want := struct{ grpb.GRIBIClient }{}
		dut.DialGRIBIFn = func(context.Context, ...grpc.DialOption) (grpb.GRIBIClient, error) {
			return want, nil
		}
		if got := gribi.New(t); got != want {
			t.Errorf("New(t) got %v, want %v", got, want)
		}
		if got := gribi.Default(t); got != want {
			t.Errorf("Default(t) got %v, want %v", got, want)
		}
	})
}

func TestP4RT(t *testing.T) {
	p4rt := NewDUTAPIs(dut).P4RT()

	t.Run("error", func(t *testing.T) {
		wantErr := "bad p4rt"
		dut.DialP4RTFn = func(context.Context, ...grpc.DialOption) (p4pb.P4RuntimeClient, error) {
			return nil, errors.New(wantErr)
		}
		gotErr := testt.ExpectFatal(t, func(t testing.TB) {
			p4rt.New(t)
		})
		if !strings.Contains(gotErr, wantErr) {
			t.Errorf("New(t) got err %v, want %v", gotErr, wantErr)
		}
		gotErr = testt.ExpectFatal(t, func(t testing.TB) {
			p4rt.Default(t)
		})
		if !strings.Contains(gotErr, wantErr) {
			t.Errorf("Default(t) got err %v, want %v", gotErr, wantErr)
		}
	})

	t.Run("success", func(t *testing.T) {
		want := struct{ p4pb.P4RuntimeClient }{}
		dut.DialP4RTFn = func(context.Context, ...grpc.DialOption) (p4pb.P4RuntimeClient, error) {
			return want, nil
		}
		if got := p4rt.New(t); got != want {
			t.Errorf("New(t) got %v, want %v", got, want)
		}
		if got := p4rt.Default(t); got != want {
			t.Errorf("Default(t) got %v, want %v", got, want)
		}
	})
}

func TestCLI(t *testing.T) {
	apis := NewDUTAPIs(dut)

	t.Run("error", func(t *testing.T) {
		wantErr := "bad cli"
		dut.DialCLIFn = func(context.Context) (binding.StreamClient, error) {
			return nil, errors.New(wantErr)
		}
		gotErr := testt.ExpectFatal(t, func(t testing.TB) {
			apis.CLI(t)
		})
		if !strings.Contains(gotErr, wantErr) {
			t.Errorf("CLI(t) got err %v, want %v", gotErr, wantErr)
		}
	})

	t.Run("success", func(t *testing.T) {
		want := struct{ binding.StreamClient }{}
		dut.DialCLIFn = func(context.Context) (binding.StreamClient, error) {
			return want, nil
		}
		if got := apis.CLI(t); want != got {
			t.Errorf("CLI(t) got %v, want %v", got, want)
		}
	})
}

func TestConsole(t *testing.T) {
	apis := NewDUTAPIs(dut)

	t.Run("error", func(t *testing.T) {
		wantErr := "bad cli"
		dut.DialConsoleFn = func(context.Context) (binding.StreamClient, error) {
			return nil, errors.New(wantErr)
		}
		gotErr := testt.ExpectFatal(t, func(t testing.TB) {
			apis.Console(t)
		})
		if !strings.Contains(gotErr, wantErr) {
			t.Errorf("Console(t) got err %v, want %v", gotErr, wantErr)
		}
	})

	t.Run("success", func(t *testing.T) {
		want := struct{ binding.StreamClient }{}
		dut.DialConsoleFn = func(context.Context) (binding.StreamClient, error) {
			return want, nil
		}
		if got := apis.Console(t); want != got {
			t.Errorf("Console(t) got %v, want %v", got, want)
		}
	})
}
