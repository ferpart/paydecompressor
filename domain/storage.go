package domain

import "encoding/json"

type IStorageService interface {
	PutFile(string, File) (json.Marshaler, error)
	ListFiles(string) (json.Marshaler, error)
}
