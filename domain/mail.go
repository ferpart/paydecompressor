package domain

import (
	"encoding/json"
)

type IMailService interface {
	GetMessageList(string, string) (json.Marshaler, error)
	GetAttachments(string, string) ([]*Attachment, error)
}

// Attachment stores the ids and data of a gmail attachment.
type Attachment struct {
	ID         string
	UserID     string
	MessageID  string
	Date       string
	DataBase64 string
	Data       []byte
}
