package internal

import (
	"fmt"
	"log"
	"os"

	"github.com/syrshax/invoice-go-v2/models"
)

func ProcessUpload(f models.FormValues, jobID string) error {
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

	err = GenerateHTMLInvoices(csv, f)
	if err != nil {
		return err
	}

	err = ConvertHTMLToPDF("invoices")
	if err != nil {
		return err
	}

	zipfilepath := "invoices_" + f.FileName + ".zip"
	err = GenerateZip("pdfs", zipfilepath)
	if err != nil {
		return err
	}
	UpdateJobPath(jobID, zipfilepath)

	fmt.Println(csv)

	return nil
}
