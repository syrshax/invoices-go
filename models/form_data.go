package models

type FormValues struct {
	Concept           string
	Date              string
	InvoiceNumber     int
	Quantity          float64
	Rate              float64
	Taxes             float64
	TypeContract      string
	UploadCsvTempPath string
}
