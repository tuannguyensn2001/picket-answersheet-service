// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: answer_sheet.proto

/*
Package answersheetpb is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package answersheetpb

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_AnswerSheetService_StartDoTest_0(ctx context.Context, marshaler runtime.Marshaler, client AnswerSheetServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq StartDoTestRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["version"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "version")
	}

	protoReq.Version, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "version", err)
	}

	msg, err := client.StartDoTest(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_AnswerSheetService_StartDoTest_0(ctx context.Context, marshaler runtime.Marshaler, server AnswerSheetServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq StartDoTestRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["version"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "version")
	}

	protoReq.Version, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "version", err)
	}

	msg, err := server.StartDoTest(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterAnswerSheetServiceHandlerServer registers the http handlers for service AnswerSheetService to "mux".
// UnaryRPC     :call AnswerSheetServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterAnswerSheetServiceHandlerFromEndpoint instead.
func RegisterAnswerSheetServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server AnswerSheetServiceServer) error {

	mux.Handle("POST", pattern_AnswerSheetService_StartDoTest_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/answer_sheet.AnswerSheetService/StartDoTest", runtime.WithHTTPPathPattern("/api/{version}/answersheets/start"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_AnswerSheetService_StartDoTest_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_AnswerSheetService_StartDoTest_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterAnswerSheetServiceHandlerFromEndpoint is same as RegisterAnswerSheetServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterAnswerSheetServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterAnswerSheetServiceHandler(ctx, mux, conn)
}

// RegisterAnswerSheetServiceHandler registers the http handlers for service AnswerSheetService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterAnswerSheetServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterAnswerSheetServiceHandlerClient(ctx, mux, NewAnswerSheetServiceClient(conn))
}

// RegisterAnswerSheetServiceHandlerClient registers the http handlers for service AnswerSheetService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "AnswerSheetServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "AnswerSheetServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "AnswerSheetServiceClient" to call the correct interceptors.
func RegisterAnswerSheetServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client AnswerSheetServiceClient) error {

	mux.Handle("POST", pattern_AnswerSheetService_StartDoTest_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/answer_sheet.AnswerSheetService/StartDoTest", runtime.WithHTTPPathPattern("/api/{version}/answersheets/start"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_AnswerSheetService_StartDoTest_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_AnswerSheetService_StartDoTest_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_AnswerSheetService_StartDoTest_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 1, 0, 4, 1, 5, 1, 2, 2, 2, 3}, []string{"api", "version", "answersheets", "start"}, ""))
)

var (
	forward_AnswerSheetService_StartDoTest_0 = runtime.ForwardResponseMessage
)
