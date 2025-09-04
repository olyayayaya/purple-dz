package file

import (
	"fmt"
	"os"
	"path/filepath"
)

type JsonDb struct {
	Filename string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		Filename: name,
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.Filename)
	if err != nil {
	fmt.Println(err)
	return nil, err
}
	return data, nil
}

func (db *JsonDb) CheckFileExtension() (bool){
 	return filepath.Ext(db.Filename) == ".json"
}