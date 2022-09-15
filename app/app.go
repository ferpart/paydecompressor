package app

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

type Context struct {
	Logger        *logrus.Logger
	MailClient    *http.Client
	StorageClient *http.Client
}
