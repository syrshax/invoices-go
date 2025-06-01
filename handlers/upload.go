package handlers

import (
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/syrshax/invoice-go-v2/internal"
	"github.com/syrshax/invoice-go-v2/models"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	invoiceNumber, _ := strconv.Atoi(r.FormValue("starting-invoice-number"))
	quantity, _ := strconv.ParseFloat(r.FormValue("quantity"), 32)
	rate, _ := strconv.ParseFloat(r.FormValue("rate"), 32)
	taxes, _ := strconv.ParseFloat(r.FormValue("taxes"), 32)

	file, header, err := r.FormFile("csv-file")
	if err != nil {
		http.Error(w, "Could not get CSV file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	err = internal.CreateInternalDirectories()
	if err != nil {
		http.Error(w, "Error creating directory: "+err.Error(), http.StatusInternalServerError)
	}

	filePath := filepath.Join("uploads", header.Filename)
	err = internal.GenerateTempCSVFile(filePath, file)
	if err != nil {
		http.Error(w, "Error generating temp CSV file: "+err.Error(), http.StatusInternalServerError)
	}

	f := models.FormValues{
		Concept:           r.FormValue("concept"),
		Date:              r.FormValue("invoice-date"),
		InvoiceNumber:     invoiceNumber,
		Quantity:          quantity,
		Rate:              rate,
		Taxes:             taxes,
		TypeContract:      r.FormValue("type-contract"),
		UploadCsvTempPath: filePath, //What if multiple files same name?
	}

	go internal.ProcessUpload(f)
}
