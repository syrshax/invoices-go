package internal

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/syrshax/invoice-go-v2/models"
)

func GenerateHTMLInvoices(c []CSVRow, f models.FormValues, s string) error {
	tmpl, err := template.ParseFiles("static/template.html")
	if err != nil {
		return fmt.Errorf("could not parse invoice template: %v", err)
	}

	for i, row := range c {
		subtotalValue := f.Quantity * f.Rate
		taxesValue := subtotalValue * (f.Taxes / 100.0)
		totalValue := subtotalValue + taxesValue

		inv := models.InvoiceData{
			CompanyAddress:  "Calle Monteagudo 23, Madrid, 28010",
			CompanyName:     "Selectra",
			Concept:         f.Concept,
			CustomerAddress: row.Address,
			CustomerID:      row.NationalID,
			CustomerName:    row.Name,
			InvoiceDate:     f.Date,
			InvoiceNumber:   fmt.Sprintf("%s %04d", f.TypeContract, f.InvoiceNumber+i),
			Quantity:        f.Quantity,
			Rate:            fmt.Sprintf("%.2f", f.Rate),
			SubTotal:        fmt.Sprintf("%.2f", subtotalValue),
			TaxesPercentage: f.Taxes,
			TaxesTotal:      fmt.Sprintf("%.2f", taxesValue),
			Total:           fmt.Sprintf("%.2f", totalValue),
			TypeContract:    f.TypeContract,
		}
		filename := fmt.Sprintf("invoice_%04d_%s.html", f.InvoiceNumber+i, sanitizeFilename(row.Name))
		filepath := filepath.Join("invoices", s+"_invoices", filename)

		file, err := os.Create(filepath)
		if err != nil {
			fmt.Printf("Warning: Could not create invoice file %s: %v\n", filename, err)
			continue
		}

		err = tmpl.Execute(file, inv)
		file.Close()

		if err != nil {
			fmt.Printf("Warning: Could not generate invoice %s: %v\n", filename, err)
			continue
		}

		fmt.Printf("Generated invoice: %s\n", filename)
	}

	return nil
}

func sanitizeFilename(name string) string {
	result := ""
	for _, char := range name {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			result += string(char)
		} else if char == ' ' || char == '-' || char == '_' {
			result += string(char)
		}
	}
	result = strings.ReplaceAll(result, " ", "_")
	if len(result) > 50 {
		result = result[:50]
	}
	return result
}
