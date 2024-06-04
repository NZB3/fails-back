package storage

import (
	"io"
	"log"
	"os"
	"strconv"
)

type FileWriter struct {
	Path string
}

func NewCounterStorage(path string) *FileWriter {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Println("file does not exist")
		_, err = os.Create(path)
		if err != nil {
			log.Fatal(err)
		}
	}
	return &FileWriter{Path: path}
}

func (fw FileWriter) Update(value int) {
	err := fw.WriteValue(value)
	if err != nil {
		log.Printf("ERROR write value to file: %s\n", err)
	}
}

func (fw FileWriter) WriteValue(value int) error {
	f, err := os.Create(fw.Path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(strconv.Itoa(value))
	return err
}

func (fw FileWriter) ReadValue() (int, error) {
	f, err := os.Open(fw.Path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	buf, err := io.ReadAll(f)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(string(buf))
}
