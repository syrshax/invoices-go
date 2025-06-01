package internal

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func GenerateZip(pdfdir string, zipfilename string) error {
	zipFile, err := os.Create(zipfilename)
	if err != nil {
		return fmt.Errorf("could not create ZIP file: %v", err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	pdfFiles, err := filepath.Glob(filepath.Join(pdfdir, "*.pdf"))
	if err != nil {
		return fmt.Errorf("could not read PDF files: %v", err)
	}

	if len(pdfFiles) == 0 {
		return fmt.Errorf("no PDF files found in %s", pdfdir)
	}

	fmt.Printf("Adding %d PDF files to ZIP...\n", len(pdfFiles))

	for _, pdfFile := range pdfFiles {
		err := addFileToZip(zipWriter, pdfFile)
		if err != nil {
			fmt.Printf("Warning: Could not add %s to ZIP: %v\n", filepath.Base(pdfFile), err)
			continue
		}
		fmt.Printf("✓ Added to ZIP: %s\n", filepath.Base(pdfFile))
	}

	return nil
}

func addFileToZip(zipWriter *zip.Writer, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = filepath.Base(filename)
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, file)
	return err
}
