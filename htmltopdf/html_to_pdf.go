package htmltopdf

import (
	"bytes"
	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"io"
	"log"
)

func GenerateHtmlToPdf(html io.Reader) (*bytes.Buffer, error) {
	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	// Set global options
	// Set options for this page
	pdfg.Dpi.Set(600)
	pdfg.NoCollate.Set(false)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.MarginBottom.Set(40)

	page := wkhtmltopdf.NewPageReader(html)

	// Add to document
	pdfg.AddPage(page)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		return nil, err
	}

	return pdfg.Buffer(), nil
}
