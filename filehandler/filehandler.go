package filehandler

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"todocli/contract"
)

const (
	DefaultStoragePath = "repo"
)

type FileHandler struct {
	storagePath string
}

func New(sPath string) (contract.StorageReaderWriter, error) {
	if sPath == "" {
		sPath = DefaultStoragePath
		os.MkdirAll(DefaultStoragePath, 0755)
	}

	if info, err := os.Stat(sPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("storage path is not exists")
	} else if !info.IsDir() {
		return nil, fmt.Errorf("storage path is not a directory")
	}

	return &FileHandler{
		storagePath: sPath,
	}, nil
}

func (f *FileHandler) Read(entityName string, entities any) error {
	file, err := f.getFile(entityName)
	if err != nil {
		return fmt.Errorf("error while reading file: %w", err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Scan()
	data := fileScanner.Text()

	json.Unmarshal([]byte(data), entities)

	return nil
}

func (f *FileHandler) Write(entityName string, entities any) error {
	data, err := json.Marshal(entities)
	if err != nil {
		return fmt.Errorf("error while writing data: %w", err)
	}

	file, err := f.getFile(entityName)
	if err != nil {
		return fmt.Errorf("error while writing data: %w", err)
	}

	file.Truncate(0)
	file.Seek(0, 0)
	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("error while writing data: %w", err)
	}

	return nil
}

func (f *FileHandler) getFile(fileName string) (*os.File, error) {
	file, err := os.OpenFile(filepath.Join(f.storagePath, fileName)+".json", os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return nil, fmt.Errorf("error while opening file: %w", err)
	}

	return file, nil
}
