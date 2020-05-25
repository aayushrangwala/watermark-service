package transport

import (
	"context"
	"encoding/json"

	"github.com/aayushrangwala/watermark-service/api/v1/pb"
	"github.com/aayushrangwala/watermark-service/pkg/watermark/endpoints"

	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	status    grpctransport.Handler
	serviceStatus grpctransport.Handler
	addDocument grpctransport.Handler
	get grpctransport.Handler
	watermark grpctransport.Handler
}

func NewGRPCServer(ep endpoints.Set) pb.WatermarkServer {
	return &grpcServer{
		status: grpctransport.NewServer(
			ep.StatusEndpoint,
			decodeGRPCStatusRequest,
			decodeGRPCStatusResponse,
		),
	}
}

func (g *grpcServer) Get(ctx context.Context, r *pb.GetRequest) (*pb.GetReply, error) {
	return nil, nil
}

func (g *grpcServer) ServiceStatus(ctx context.Context, r *pb.ServiceStatusRequest) (*pb.ServiceStatusReply, error) {
	return nil, nil
}

func (g *grpcServer) AddDocument(ctx context.Context, r *pb.AddDocumentRequest) (*pb.AddDocumentReply, error) {
	return nil, nil
}

func (g *grpcServer) Status(ctx context.Context, r *pb.StatusRequest) (*pb.StatusReply, error) {
	return nil, nil
}

func (g *grpcServer) Watermark(ctx context.Context, r *pb.WatermarkRequest) (*pb.WatermarkReply, error) {
	return nil, nil
}

func decodeGRPCGetRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return endpoints.GetRequest{}, nil
}

func decodeGRPCStatusRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	var req endpoints.StatusRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGRPCWatermarkRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	var req endpoints.WatermarkRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGRPCAddDocumentRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	var req endpoints.AddDocumentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGRPCServiceStatusRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	var req endpoints.ServiceStatusRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}


func decodeGRPCGetResponse(ctx context.Context, grpcReply interface{}) (interface{}, error) {
	return endpoints.GetRequest{}, nil
}

func decodeGRPCStatusResponse(ctx context.Context, grpcReply interface{}) (interface{}, error) {
	var req endpoints.StatusRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGRPCWatermarkResponse(ctx context.Context, grpcReply interface{}) (interface{}, error) {
	var req endpoints.WatermarkRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGRPCAddDocumentResponse(ctx context.Context, grpcReply interface{}) (interface{}, error) {
	var req endpoints.AddDocumentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGRPCServiceStatusResponse(ctx context.Context, grpcReply interface{}) (interface{}, error) {
	var req endpoints.ServiceStatusRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
