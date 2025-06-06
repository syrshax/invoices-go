package internal

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/syrshax/invoice-go-v2/models"
)

func ProcessUpload(f models.FormValues, jobID string) error {
	//We delete the uploaded temp file csv...
	defer func() {
		err := os.Remove(f.UploadCsvTempPath)
		if err != nil {
			log.Printf("Couldnt delete file: %v \n", f.UploadCsvTempPath)
		}

	}()

	csv, err := ReadCSV(f.UploadCsvTempPath)
	if err != nil {
		return err
	}

	err = GenerateHTMLInvoices(csv, f, jobID)
	if err != nil {
		return err
	}

	err = ConvertHTMLToPDF(jobID)
	if err != nil {
		return err
	}

	pdfDirectory := filepath.Join("pdfs", jobID+"_pdfs")
	zipDir := path.Join("zipfiles", jobID+"_zipfiles")
	err = GenerateZip(pdfDirectory, zipDir, jobID)
	if err != nil {
		return err
	}
	UpdateJobPath(jobID, zipDir)

	for _, c := range csv {
		fmt.Println(c)
	}

	return nil
}
