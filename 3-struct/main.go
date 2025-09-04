package main

import (
	"3-struct/file"
	"3-struct/storage"
)



func main() {
	db := file.NewJsonDb("bins.json")

	st := storage.NewJsonStorage(db)
}