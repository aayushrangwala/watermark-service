package watermark

import (
	"context"
	"net/http"

	"github.com/aayushrangwala/watermark-service/internal"

	"github.com/lithammer/shortuuid/v3"
	log "github.com/sirupsen/logrus"
)

type watermarkService struct {}

func NewService() Service {return &watermarkService{}}

func (w *watermarkService) Get(ctx context.Context, filters ...internal.Filter) ([]internal.Document, error) {
	// query the database using the filters and return the list of documents
	// return error if the filter (key) is invalid and also return error if no item found
	return []internal.Document{}, nil
}

func (w *watermarkService) Status(ctx context.Context, ticketID string) (internal.Status, error) {
	// query database using the ticketID and return the document info
	// return err if the ticketID is invalid or no Document exists for that ticketID
	return internal.InProgress, nil
}

func (w *watermarkService) Watermark(ctx context.Context, ticketID, mark string) (int, error) {
	// update the database entry with watermark field as non empty
	// first check if the watermark status is not already in InProgress, Started or Finished state
	// If yes, then return invalid request
	// return error if no item found using the ticketID
	return http.StatusOK, nil
}

func (w *watermarkService) AddDocument(ctx context.Context, doc internal.Document) (string, error) {
	// add the document entry in the database by calling the database service
	// return error if the doc is invalid and/or the database invalid entry error
	newTicketID := shortuuid.New()
	return newTicketID, nil
}


func (w *watermarkService) ServiceStatus(ctx context.Context) (int, error) {
	log.Infof("Checking the Service health...")
	return http.StatusOK, nil
}