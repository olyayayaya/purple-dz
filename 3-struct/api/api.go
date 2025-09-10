package api

import (
	"3-struct/config"
	"3-struct/storage"
)

type API struct {
	st storage.Storage
	config config.Config
}

func NewAPI(st storage.Storage, config config.Config) *API {
	return &API{
		st: st,
		config: config,
	}
}