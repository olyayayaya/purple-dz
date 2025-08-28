package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func ReadFile(fileName string) ([]byte, error) {
if filepath.Ext(fileName) == ".json" {
	data, err := os.ReadFile(fileName)
	if err != nil {
	fmt.Println(err)
	return nil, err
}
	return data, nil
} else {
	return nil, errors.New("неверное расширение файла")
}
}