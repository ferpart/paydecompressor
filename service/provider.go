package service

import (
	"github.com/ferpart/paydecompress/app"
	"github.com/ferpart/paydecompress/domain"
	"github.com/ferpart/paydecompress/service/gdrive"
	"github.com/ferpart/paydecompress/service/gmail"
	"github.com/ferpart/paydecompress/service/zip"
)

// Provider implements all the services used by the application
type Provider struct {
	MailService    domain.IMailService
	StorageService domain.IStorageService
	ZipService     domain.IZipService
}

// New returns a new Provider using the application context configuration.
func New(appContext *app.Context) *Provider {
	return &Provider{
		MailService:    gmail.New(appContext.MailClient, appContext.Logger),
		StorageService: gdrive.New(appContext.StorageClient, appContext.Logger),
		ZipService:     zip.New(appContext.Logger),
	}
}
