package internal

import "os"

// It creates uploads, pdfs and invoices directories empty ready for use.
func CreateInternalDirectories() error {
	err := os.MkdirAll("uploads", 0775)
	if err != nil {
		return err
	}
	err = os.MkdirAll("pdfs", 0775)
	if err != nil {
		return err
	}

	err = os.MkdirAll("invoices", 0775)
	if err != nil {
		return err
	}
	return nil
}

func CleaningFiles() error {
	return nil
}
