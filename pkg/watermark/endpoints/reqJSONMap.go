package endpoints

import "github.com/aayushrangwala/watermark-service/internal"

type GetRequest struct {
	Filters []internal.Filter `json:"filters,omitempty"`
}

type GetResponse struct {
	Documents []internal.Document `json:"documents"`
	Err       string              `json:"err,omitempty"`
}

type StatusRequest struct {
	TicketID string `json:"ticketID"`
}

type StatusResponse struct {
	Status internal.Status `json:"status"`
	Err    string          `json:"err,omitempty"`
}

type WatermarkRequest struct {
	TicketID string `json:"ticketID"`
	Mark     string `json:"mark"`
}

type WatermarkResponse struct {
	Code int    `json:"code"`
	Err  string `json:"err"`
}

type AddDocumentRequest struct {
	Document *internal.Document `json:"document"`
}

type AddDocumentResponse struct {
	TicketID string `json:"ticketID"`
	Err      string `json:"err,omitempty"`
}

type ServiceStatusRequest struct{}

type ServiceStatusResponse struct {
	Code int    `json:"status"`
	Err  string `json:"err,omitempty"`
}
