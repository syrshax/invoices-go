package models

import ()

type InvoiceData struct {
	Concept         string
	CompanyAddress  string
	CompanyName     string
	CustomerAddress string
	CustomerID      string
	CustomerName    string
	InvoiceDate     string
	InvoiceNumber   string
	Quantity        float64
	Rate            float64
	SubTotal        string
	TaxesPercentage float64
	TaxesTotal      string
	Total           string
	TypeContract    string
}
