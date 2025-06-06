package internal

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

func ConvertHTMLToPDF(jobID string) error {
	_, err := exec.LookPath("wkhtmltopdf")
	if err != nil {
		return fmt.Errorf("wkhtmltopdf not found. Please install it first: %v", err)
	}

	files, err := filepath.Glob(filepath.Join("invoices", jobID+"_invoices", "*.html"))
	if err != nil {
		return fmt.Errorf("could not read HTML files: %v", err)
	}

	fmt.Println("FILES!!!!", files)

	if len(files) == 0 {
		return fmt.Errorf("no HTML files found in %s", jobID+"_invoices")
	}

	headerTemplatePath, err := filepath.Abs(filepath.Join("static", "header.html"))
	if err != nil {
		return fmt.Errorf("error resolving header template path: %v", err)
	}
	footerTemplatePath, err := filepath.Abs(filepath.Join("static", "footer.html"))
	if err != nil {
		return fmt.Errorf("error resolving footer template path: %v", err)
	}

	var wg sync.WaitGroup

	for _, htmlFile := range files {
		fmt.Println(htmlFile, files)
		wg.Add(1)

		go func(htmlToConvert string) {
			defer wg.Done()
			baseName := strings.TrimSuffix(filepath.Base(htmlToConvert), ".html")
			pdfFile := filepath.Join("pdfs", jobID+"_pdfs", baseName+".pdf")

			fmt.Println("BASE NAME", baseName)
			fmt.Println("PDF FILE", pdfFile)

			cmd := exec.Command("wkhtmltopdf",
				"--page-size", "A4",
				"--margin-top", "1.5in",
				"--margin-right", "0.75in",
				"--margin-bottom", "1.5in",
				"--margin-left", "0.75in",
				"--encoding", "UTF-8",
				"--no-outline",
				"--header-html", headerTemplatePath,
				"--footer-html", footerTemplatePath,
				"--enable-smart-shrinking",
				"--enable-local-file-access",
				htmlToConvert,
				pdfFile,
			)

			output, err := cmd.CombinedOutput()
			if err != nil {
				log.Printf("Warning: Failed to convert %s to PDF: %v\nOutput: %s\n",
					filepath.Base(htmlToConvert), err, string(output))
				return
			}

			log.Printf("Generated PDF: %s\n", filepath.Base(pdfFile))

		}(htmlFile)
	}
	wg.Wait()

	log.Println("All PDF conversions complete.")
	return nil
}
