package endpoints

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/aayushrangwala/watermark-service/internal"
	"github.com/aayushrangwala/watermark-service/pkg/database"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

type Set struct {
	GetEndpoint           endpoint.Endpoint
	AddEndpoint           endpoint.Endpoint
	UpdateEndpoint        endpoint.Endpoint
	RemoveEndpoint        endpoint.Endpoint
	ServiceStatusEndpoint endpoint.Endpoint
}

func NewEndpointSet(svc database.Service) Set {
	return Set{
		GetEndpoint:           MakeGetEndpoint(svc),
		AddEndpoint:           MakeAddEndpoint(svc),
		UpdateEndpoint:        MakeUpdateEndpoint(svc),
		RemoveEndpoint:        MakeRemoveEndpoint(svc),
		ServiceStatusEndpoint: MakeServiceStatusEndpoint(svc),
	}
}

func MakeGetEndpoint(svc database.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		docs, err := svc.Get(ctx, req.Filters...)
		if err != nil {
			return GetResponse{docs, err.Error()}, nil
		}
		return GetResponse{docs, ""}, nil
	}
}

func MakeUpdateEndpoint(svc database.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		code, err := svc.Update(ctx, req.TicketID, req.Document)
		if err != nil {
			return UpdateResponse{Code: code, Err: err.Error()}, nil
		}
		return UpdateResponse{Code: code, Err: ""}, nil
	}
}

func MakeAddEndpoint(svc database.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRequest)
		ticketID, err := svc.Add(ctx, req.Document)
		if err != nil {
			return AddResponse{TicketID: ticketID, Err: err.Error()}, nil
		}
		return AddResponse{TicketID: ticketID, Err: ""}, nil
	}
}

func MakeRemoveEndpoint(svc database.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RemoveRequest)
		code, err := svc.Remove(ctx, req.TicketID)
		if err != nil {
			return RemoveResponse{Code: code, Err: err.Error()}, nil
		}
		return RemoveResponse{Code: code, Err: ""}, nil
	}
}

func MakeServiceStatusEndpoint(svc database.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(ServiceStatusRequest)
		code, err := svc.ServiceStatus(ctx)
		if err != nil {
			return ServiceStatusResponse{Code: code, Err: err.Error()}, nil
		}
		return ServiceStatusResponse{Code: code, Err: ""}, nil
	}
}

func (s *Set) Get(ctx context.Context, filters ...internal.Filter) ([]internal.Document, error) {
	resp, err := s.GetEndpoint(ctx, GetRequest{Filters: filters})
	if err != nil {
		return []internal.Document{}, err
	}
	getResp := resp.(GetResponse)
	if getResp.Err != "" {
		return []internal.Document{}, errors.New(getResp.Err)
	}
	return getResp.Documents, nil
}

func (s *Set) ServiceStatus(ctx context.Context) (int, error) {
	resp, err := s.ServiceStatusEndpoint(ctx, ServiceStatusRequest{})
	svcStatusResp := resp.(ServiceStatusResponse)
	if err != nil {
		return svcStatusResp.Code, err
	}
	if svcStatusResp.Err != "" {
		return svcStatusResp.Code, errors.New(svcStatusResp.Err)
	}
	return svcStatusResp.Code, nil
}

func (s *Set) Add(ctx context.Context, doc *internal.Document) (string, error) {
	resp, err := s.AddEndpoint(ctx, AddRequest{Document: doc})
	if err != nil {
		return "", err
	}
	adResp := resp.(AddResponse)
	if adResp.Err != "" {
		return "", errors.New(adResp.Err)
	}
	return adResp.TicketID, nil
}

func (s *Set) Update(ctx context.Context, ticketID string, doc *internal.Document) (int, error) {
	resp, err := s.UpdateEndpoint(ctx, UpdateRequest{TicketID: ticketID, Document: doc})
	if err != nil {
		return http.StatusBadRequest, err
	}
	stsResp := resp.(UpdateResponse)
	if stsResp.Err != "" {
		return http.StatusConflict, errors.New(stsResp.Err)
	}
	return http.StatusOK, nil
}

func (s *Set) Remove(ctx context.Context, ticketID string) (int, error) {
	resp, err := s.RemoveEndpoint(ctx, RemoveRequest{TicketID: ticketID})
	wmResp := resp.(RemoveResponse)
	if err != nil {
		return wmResp.Code, err
	}
	if wmResp.Err != "" {
		return wmResp.Code, errors.New(wmResp.Err)
	}
	return wmResp.Code, nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
