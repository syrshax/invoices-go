package internal

import (
	"os"
	"path/filepath"
)

func CreateInternalDirectories(s string) error {
	err := os.MkdirAll("uploads", 0775)
	if err != nil {
		return err
	}
	err = os.MkdirAll(filepath.Join("pdfs", s+"_pdfs"), 0775)
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Join("invoices", s+"_invoices"), 0775)
	if err != nil {
		return err
	}

	err = os.MkdirAll("zipfiles", 0775)
	if err != nil {
		return err
	}
	return nil
}

func CleaningFiles(s string) error {
	err := os.RemoveAll(filepath.Join("pdfs", s+"_pdfs"))
	if err != nil {
		return err
	}
	err = os.RemoveAll(filepath.Join("invoices", s+"_invoices"))
	if err != nil {
		return err
	}
	err = os.RemoveAll(filepath.Join("zipfiles", s+"_zipfiles"))
	if err != nil {
		return err
	}

	return nil
}
