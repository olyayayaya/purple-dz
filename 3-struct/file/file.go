package file

import (
	"fmt"
	"os"
	"path/filepath"
)

func ReadFile(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
	fmt.Println(err)
	return nil, err
}
	return data, nil
}

func CheckFileExtension(fileName string) (bool){
 	return filepath.Ext(fileName) == ".json"
}