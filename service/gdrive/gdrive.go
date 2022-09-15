package gdrive

import (
	"context"
	"encoding/json"
	"github.com/ferpart/paydecompress/domain"
	"net/http"

	"github.com/sirupsen/logrus"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type Service struct {
	gdService *drive.Service
	log       *logrus.Logger
}

// New creates a new drive Service. A fatal log will be sent if service is unable
// to start.
func New(gClient *http.Client, logger *logrus.Logger) *Service {
	srv, err := drive.NewService(context.TODO(), option.WithHTTPClient(gClient))
	if err != nil {
		logger.Fatal(err.Error())
	}
	return &Service{
		gdService: srv,
		log:       logger,
	}
}

// ListFiles returns a list of all files with a given query. Will return an error
// if query is invalid.
func (s *Service) ListFiles(query string) (json.Marshaler, error) {
	return s.gdService.Files.List().Q(query).Do()
}

// PutFile puts a domain.File into the given file path. Will return an error for
// any non 2xx status code.
func (s *Service) PutFile(filePath string, file domain.File) (json.Marshaler, error) {
	return s.gdService.Files.Create(
		&drive.File{
			Name: filePath + file.Name,
		},
	).Media(file.Body).Do()
}
