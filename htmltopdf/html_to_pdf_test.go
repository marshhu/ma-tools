package htmltopdf

import (
	"io/ioutil"
	"strings"
	"testing"
)

func Test_GenerateHtmlToPdf(t *testing.T) {
	html, err := ioutil.ReadFile("./template.html")
	if err != nil {
		panic(err)
	}
	reader := strings.NewReader(string(html))
	buffer, err := GenerateHtmlToPdf(reader)
	if err != nil {
		t.FailNow()
	}
	err = ioutil.WriteFile("./result.pdf", buffer.Bytes(), 0644)
	if err != nil {
		t.FailNow()
	}
}
