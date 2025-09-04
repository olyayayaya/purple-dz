package storage

import (
	"3-struct/bins"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Storage interface {
	Write(*bins.BinList,string)
	Read()(*bins.BinList)
}

type JsonStorage struct {
	db bins.Db
}

func NewJsonStorage(db bins.Db) *JsonStorage {
	return &JsonStorage{db: db}
}

func (s *JsonStorage) Write(content *bins.BinList, name string) {
jsonData, err := json.Marshal(content)
if err != nil {
	fmt.Println(err)
}

 file, err := os.Create(name)
 if err != nil {
	fmt.Println(err)
 }

_, err = file.Write(jsonData)
 if err != nil {
	file.Close()
	fmt.Println(err)
	return
 }
 	fmt.Println("запись успешнa")
}


func (s *JsonStorage) Read() *bins.BinList {
		file, err := s.db.Read()
	if err != nil {	
		return &bins.BinList {
		Bins: []bins.Bin{}, 
		UpdatedAt: time.Now(),
	}
}

var binlist bins.BinList 
err = json.Unmarshal(file, &binlist)
	if err != nil {	
		fmt.Println(err.Error())
		return &bins.BinList {
		Bins: []bins.Bin{}, 
		UpdatedAt: time.Now(),
	}
	}
return &binlist
}