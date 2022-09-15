package gmail

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/api/option"
	"net/http"

	"github.com/ferpart/paydecompress/domain"
	"github.com/sirupsen/logrus"

	"google.golang.org/api/gmail/v1"
)

type Service struct {
	gmService *gmail.Service
	log       *logrus.Logger
}

// New creates a new gmail Service. A fatal log will be sent if service is unable
// to start.
func New(gClient *http.Client, logger *logrus.Logger) *Service {
	srv, err := gmail.NewService(context.TODO(), option.WithHTTPClient(gClient))
	if err != nil {
		logger.Fatal(err.Error())
	}
	return &Service{
		gmService: srv,
		log:       logger,
	}
}

// GetMessageList takes a userID and a query to obtain a list of gmail messages.
// Returns an error if the query fails.
func (s *Service) GetMessageList(userID, query string) (json.Marshaler, error) {
	return s.gmService.Users.Messages.List(userID).Q(query).Do()
}

// GetAttachments takes a userID and a query to obtain a list of Attachment.
// Returns an error if the message query fails, or no attachments are found.
func (s *Service) GetAttachments(userID, query string) ([]*domain.Attachment, error) {
	resp, err := s.GetMessageList(userID, query)
	if err != nil {
		return nil, err
	}

	messageList, ok := resp.(*gmail.ListMessagesResponse)
	if !ok {
		return nil, fmt.Errorf(
			"failed to obtain messages from '%s' using query '%s'", userID, query,
		)
	}

	attachments := make([]*domain.Attachment, 0)
	for _, msg := range messageList.Messages {
		for _, part := range msg.Payload.Parts {
			if part.Filename != "" {
				attachment := &domain.Attachment{
					ID:        part.Body.AttachmentId,
					MessageID: msg.Id,
					UserID:    userID,
					Date:      msg.Header.Get("Date"),
				}
				if part.Body.Data != "" {
					attachment.DataBase64 = part.Body.Data
				} else {
					_, err := s.getAttachmentData(attachment)
					if err != nil {
						s.log.Error(err.Error())
						continue
					}
				}

				attachments = append(attachments, attachment)
			}
		}
	}

	if len(attachments) == 0 {
		return nil, fmt.Errorf(
			"no attachments found for query \"%s\" on user %s", query, userID,
		)
	}
	return attachments, nil
}

func (s *Service) getAttachmentData(attachment *domain.Attachment) (*domain.Attachment, error) {
	gAttachment, err := s.gmService.Users.Messages.Attachments.Get(
		attachment.UserID,
		attachment.MessageID,
		attachment.ID).Do()
	if err != nil {
		return attachment, err
	}

	attachment.DataBase64 = gAttachment.Data

	return attachment, nil
}
