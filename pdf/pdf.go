package pdf

import (
	"bytes"
	"html/template"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

// define the path of the wkhtmltopdf
const path = "/usr/bin/wkhtmltopdf"

func Gerar(file string, pdfFile string) error {
	invoice := map[string]interface{}{
		"name": "teste nome empresa",
	}
	var processed bytes.Buffer

	tmpl, err := template.ParseFiles(file)

	tmpl.Execute(&processed, invoice)
	// read the html content
	//html, err := ioutil.ReadFile(file)
	//if err != nil {
	//	return err
	//}

	// set the predefined path in the wkhtmltopdf's global state
	wkhtmltopdf.SetPath(path)

	//template.Must(template.New("transform").Parse())

	// create a new page based on the HTML
	page := wkhtmltopdf.NewPageReader(bytes.NewReader(processed.Bytes()))
	page.NoBackground.Set(true)
	page.DisableExternalLinks.Set(false)

	// create a new instance of the PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return err
	}

	// add page to the PDF generator
	pdfg.AddPage(page)

	// set dpi of the content
	pdfg.Dpi.Set(350)

	// set margins to zero at all direction
	pdfg.MarginBottom.Set(0)
	pdfg.MarginTop.Set(0)
	pdfg.MarginLeft.Set(0)
	pdfg.MarginRight.Set(0)

	// create the exact pdf
	err = pdfg.Create()
	if err != nil {
		return err
	}

	// write it into a file
	err = pdfg.WriteFile(pdfFile)
	if err != nil {
		return err
	}

	return nil
}
