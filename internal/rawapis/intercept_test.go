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

package rawapis

import (
	"golang.org/x/net/context"
	"fmt"
	"io"
	"strings"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/openconfig/ondatra/binding/grpcutil/testservice"

	tgrpcpb "github.com/openconfig/ondatra/binding/grpcutil/testservice/gen"
	tpb "github.com/openconfig/ondatra/binding/grpcutil/testservice/gen"
)

func TestAnnotateErrors(t *testing.T) {
	errSrv := new(errServer)
	ctx := context.Background()
	serv := testservice.Start(ctx, t, errSrv,
		withUnaryAnnotateErrors(),
		withStreamAnnotateErrors())
	defer serv.Stop()

	t.Run("unary success", func(t *testing.T) {
		errSrv.unaryErr = nil
		if _, gotErr := serv.SendUnary(ctx, new(tpb.TestRequest)); gotErr != nil {
			t.Errorf("withUnaryAnnotateErrors() got error %v, want no error", gotErr)
		}
	})

	t.Run("unary error", func(t *testing.T) {
		const (
			wantCode   = codes.NotFound
			wantErrMsg = "unary error"
			wantReqMsg = "unary msg"
		)
		errSrv.unaryErr = status.Error(wantCode, wantErrMsg)
		_, gotErr := serv.SendUnary(ctx, &tpb.TestRequest{Message: wantReqMsg})
		if diff := grpcErrDiff(t, gotErr, wantCode, wantErrMsg, wantReqMsg); diff != "" {
			t.Errorf("withUnaryAnnotateErrors() got grpc error diff:\n%s", diff)
		}
	})

	t.Run("unary unknown error", func(t *testing.T) {
		const (
			wantCode   = codes.Unknown
			wantErrMsg = "unknown error"
			wantReqMsg = "unknown msg"
		)
		errSrv.unaryErr = fmt.Errorf(wantErrMsg)
		_, gotErr := serv.SendUnary(ctx, &tpb.TestRequest{Message: wantReqMsg})
		if diff := grpcErrDiff(t, gotErr, wantCode, wantErrMsg, wantReqMsg); diff != "" {
			t.Errorf("withUnaryAnnotateErrors() got grpc error diff:\n%s", diff)
		}
	})

	t.Run("stream recv eof", func(t *testing.T) {
		errSrv.streamErr = nil
		streamClient := serv.MustSendStream(ctx, t)
		streamClient.MustSend(t, "msg")
		if _, gotErr := streamClient.Recv(); gotErr != io.EOF {
			t.Errorf("withStreamAnnotateErrors() got error %v, want %v", gotErr, io.EOF)
		}
	})

	t.Run("stream recv no send", func(t *testing.T) {
		const (
			wantCode   = codes.PermissionDenied
			wantErrMsg = "stream error"
		)
		errSrv.streamErr = status.Error(wantCode, wantErrMsg)
		streamClient := serv.MustSendStream(ctx, t)
		_, gotErr := streamClient.Recv()
		if diff := grpcErrDiff(t, gotErr, wantCode, wantErrMsg); diff != "" {
			t.Errorf("withStreamAnnotateErrors() got grpc error diff:\n%s", diff)
		}
	})

	t.Run("stream recv after send", func(t *testing.T) {
		const (
			wantCode   = codes.PermissionDenied
			wantErrMsg = "stream error"
			wantReqMsg = "stream msg"
		)
		errSrv.streamErr = status.Error(wantCode, wantErrMsg)
		streamClient := serv.MustSendStream(ctx, t)
		streamClient.MustSend(t, wantReqMsg)
		_, gotErr := streamClient.Recv()
		if gotErr == nil {
			t.Fatalf("withStreamAnnotateErrors() got no error, want error, %v", gotErr)
		}
		if diff := grpcErrDiff(t, gotErr, wantCode, wantErrMsg, wantReqMsg); diff != "" {
			t.Errorf("withStreamAnnotateErrors() got grpc error diff:\n%s", diff)
		}
	})
}

type errServer struct {
	tgrpcpb.UnimplementedTestServer
	unaryErr  error
	streamErr error
}

func (s *errServer) SendUnary(context.Context, *tpb.TestRequest) (*tpb.TestResponse, error) {
	return new(tpb.TestResponse), s.unaryErr
}

func (s *errServer) SendStream(srv tgrpcpb.Test_SendStreamServer) error {
	return s.streamErr
}

func grpcErrDiff(t *testing.T, err error, wantCode codes.Code, wantSubstrs ...string) string {
	gotCode, gotMsg := codes.OK, ""
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			t.Fatalf("status.FromError failed to convert error: %v", err)
		}
		gotCode, gotMsg = st.Code(), st.Message()
	}
	var lines []string
	if gotCode != wantCode {
		lines = append(lines, fmt.Sprintf("want code %v, got %v", wantCode, gotCode))
	}
	for _, wantSub := range wantSubstrs {
		if !strings.Contains(gotMsg, wantSub) {
			lines = append(lines, fmt.Sprintf("want message containing %q, got %q", wantSub, gotMsg))
		}
	}
	return strings.Join(lines, "\n")
}
