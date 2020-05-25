package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aayushrangwala/watermark-service/internal/util"
	"github.com/aayushrangwala/watermark-service/pkg/watermark/endpoints"

	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPHandler(ep endpoints.Set) http.Handler {
	m := http.NewServeMux()

	m.Handle("/healthz", httptransport.NewServer(
		ep.ServiceStatusEndpoint,
		decodeHTTPServiceStatusRequest,
		encodeResponse,
	))
	m.Handle("/status", httptransport.NewServer(
		ep.StatusEndpoint,
		decodeHTTPStatusRequest,
		encodeResponse,
	))
	m.Handle("/addDocument", httptransport.NewServer(
		ep.AddDocumentEndpoint,
		decodeHTTPAddDocumentRequest,
		encodeResponse,
	))
	m.Handle("/get", httptransport.NewServer(
		ep.GetEndpoint,
		decodeHTTPGetRequest,
		encodeResponse,
	))
	m.Handle("/watermark", httptransport.NewServer(
		ep.WatermarkEndpoint,
		decodeHTTPWatermarkRequest,
		encodeResponse,
	))

	return m
}

func decodeHTTPGetRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return endpoints.GetRequest{}, nil
}

func decodeHTTPStatusRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.StatusRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeHTTPWatermarkRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.WatermarkRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeHTTPAddDocumentRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.AddDocumentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeHTTPServiceStatusRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.ServiceStatusRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(error); ok && e != nil {
		encodeError(ctx, e, w)
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case util.ErrUnknown:
		w.WriteHeader(http.StatusNotFound)
	case util.ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
