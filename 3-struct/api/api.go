package api


import (
	"3-struct/storage"
)

type API struct {
	st storage.Storage
}

func NewAPI(st storage.Storage) *API {
	return &API{st: st}
}