package storage

import (
	"3-struct/bins"
	"3-struct/file"
	"encoding/json"
	"fmt"
	"os"
	"time"
)


func WriteFile(content *bins.BinList, name string) {
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


func ReadFile(name string) *bins.BinList {
		file, err := file.ReadFile(name)
	if err != nil {	
		return &bins.BinList {
		Bins: []*bins.Bin{}, 
		UpdatedAt: time.Now(),
	}
}

var binlist bins.BinList 
err = json.Unmarshal(file, &binlist)
	if err != nil {	
		fmt.Println(err.Error())
		return &bins.BinList {
		Bins: []*bins.Bin{}, 
		UpdatedAt: time.Now(),
	}
	}
return &binlist
}