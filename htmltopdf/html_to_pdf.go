package htmltopdf

import (
	"fmt"
	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"log"
	"strings"
)

func ExampleNewPDFGenerator() {

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

	////// Create a new input page from an URL
	//page := wkhtmltopdf.NewPage("https://godoc.org/github.com/SebastiaanKlippert/go-wkhtmltopdf")
	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <h1 class="ql-align-center"><span class="ql-font-serif" style="color: rgb(255, 255, 255); background-color: rgb(240, 102, 102);"> I am snow example! </span></h1>
    <p><br></p>
    <p><span class="ql-font-serif">W Can a man still be brave if he's afraid? That is the only time a man can be brave. </span></p>
    <p><br></p>
    <p><strong class="ql-size-large ql-font-serif">Courage and folly is </strong>
        <strong class="ql-size-large ql-font-serif" style="color: rgb(230, 0, 0);">always</strong>
        <strong class="ql-size-large ql-font-serif"> just a fine line.</strong></p>
        <p><br></p>
        <p><u class="ql-font-serif">There is only one God, and his name is Death. And there is only one thing we say to Death: "Not today."</u></p>
        <p><br></p>
        <p><em class="ql-font-serif">Fear cuts deeper than swords.</em></p>
        <p><br></p>
        <p><br></p>
        <p><span class="ql-font-serif">Every flight begins with a fall.</span></p>
        <p><br></p>
        <p><a href="https://surmon.me/" rel="noopener noreferrer" target="_blank" class="ql-size-small ql-font-serif" style="color: rgb(230, 0, 0);"><u>A ruler who hides behind paid executioners soon forgets what death is. </u></a></p>
        <p><br></p>
        <p><span class="ql-font-serif">We are born to suffer, to suffer can make us strong.</span></p>
        <p><br></p>
        <p><span class="ql-font-serif">The things we love destroy us every time.</span></p>
       <p><img src="http://a1.att.hudong.com/62/02/01300542526392139955025309984.jpg"/></p>
</body>
</html>`
	page := wkhtmltopdf.NewPageReader(strings.NewReader(html))

	// Add to document
	pdfg.AddPage(page)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile("./simplesample.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
	// Output: Done
}
