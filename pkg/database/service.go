package database

import (
	"context"

	"github.com/aayushrangwala/watermark-service/internal"
)

type Service interface {
	Add(ctx context.Context, doc *internal.Document) (string, error)
	// Get the list of all documents
	Get(ctx context.Context, filters ...internal.Filter) ([]internal.Document, error)
	Update(ctx context.Context, ticketID string, doc *internal.Document) (int, error)
	Remove(ctx context.Context, ticketID string) (int, error)
	//Validate(ctx context.Context, doc *internal.Document) (bool, error)
	ServiceStatus(ctx context.Context) (int, error)
}
