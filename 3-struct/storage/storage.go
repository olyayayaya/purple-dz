package storage

import (
	"3-struct/bins"
	"3-struct/file"
	"encoding/json"
	"fmt"
	"os"
	"time"
)


func (bins *bins.Bin) WriteFile() (content []byte, name string) {
 file, err := os.Create(name)
 if err != nil {
	fmt.Println(err)
 }
_, err = file.Write(content)
 if err != nil {
	file.Close()
	fmt.Println(err)
	return
 }
 	fmt.Println("запись успешнa")
	return
}