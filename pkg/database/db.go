package database

import (
	"context"
	"net/http"

	"github.com/aayushrangwala/watermark-service/internal"
)

type dbService struct{}

func NewService() Service { return &dbService{} }

func (d *dbService) Add(_ context.Context, doc *internal.Document) (string, error) {
	return "", nil
}

func (d *dbService) Get(_ context.Context, filters ...internal.Filter) ([]internal.Document, error) {
	return []internal.Document{}, nil
}

func (d *dbService) Update(_ context.Context, ticketId string, doc *internal.Document) (int, error) {
	return http.StatusOK, nil
}

func (d *dbService) Remove(_ context.Context, ticketId string) (int, error) {
	return http.StatusOK, nil
}

//func (d *dbService) Validate(_ context.Context, doc *internal.Document) (bool, error) {
//	return true, nil
//}

func (d *dbService) ServiceStatus(_ context.Context) (int, error) {
	return http.StatusOK, nil
}
