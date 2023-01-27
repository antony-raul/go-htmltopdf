package main

import (
	"log"

	"github.com/antony-raul/go-htmltopdf/pdf"
)

func main() {
	if err := pdf.Gerar("examples/invoice-template-1/invoice.html", "./invoice1.pdf"); err != nil {
		log.Println(err)
	}
}
