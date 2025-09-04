package main

import (
	"3-struct/api"
	"3-struct/bins"
	"3-struct/file"
	"3-struct/storage"
)



func main() {
	db := file.NewJsonDb("bins.json")

	st := storage.NewJsonStorage(db)

	bin := bins.NewBin("1", false, "TestBin")
	binList := &bins.BinList{
		Bins:      []bins.Bin{*bin},
		UpdatedAt: bin.CreatedAt,
	}


	server := api.NewAPI(st)

}