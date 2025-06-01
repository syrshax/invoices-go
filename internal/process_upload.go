package internal

import (
	"fmt"
	"log"
	"os"

	"github.com/syrshax/invoice-go-v2/models"
)

func ProcessUpload(f models.FormValues) error {
	//We delete the uploaded temp file csv...
	defer func() {
		err := os.Remove(f.UploadCsvTempPath)
		if err != nil {
			log.Printf("Couldnt delete file: %v \n", f.UploadCsvTempPath)
		}

		// _ = os.RemoveAll("invoices")
	}()

	csv, err := ReadCSV(f.UploadCsvTempPath)
	if err != nil {
		return err
	}

	GenerateHTMLInvoices(csv, f)

	fmt.Println(csv)

	return nil
}
