package internal

import (
	"io"
	"mime/multipart"
	"os"
)

func GenerateTempCSVFile(filename string, filedata multipart.File) error {
	emptyFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	_, err = io.Copy(emptyFile, filedata)
	if err != nil {
		return err
	}
	return nil
}
