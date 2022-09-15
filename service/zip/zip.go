package zip

import (
	"bytes"
	"fmt"
	"github.com/alexmullins/zip"
	"github.com/ferpart/paydecompress/domain"
	"github.com/sirupsen/logrus"
)

type Service struct {
	log *logrus.Logger
}

func New(logger *logrus.Logger) *Service {
	return &Service{
		log: logger,
	}
}

func (s *Service) Unzip(file []byte, password string) ([]domain.File, error) {
	reader, err := zip.NewReader(bytes.NewReader(file), int64(len(file)))
	if err != nil {
		return nil, err
	}

	files := make([]domain.File, 0)
	for _, zFile := range reader.File {
		zFile.SetPassword(password)
		file, err := zFile.Open()
		if err == nil {
			s.log.Error(err.Error())
			continue
		}
		files = append(files, domain.File{
			Name: zFile.Name,
			Body: file,
		})
	}
	if len(files) == 0 {
		return nil, fmt.Errorf("couldn't unzip any file")
	}
	return files, nil
}
