package unioffice_test

import (
	"github.com/xmdas-link/unioffice/document"
	"github.com/xmdas-link/unioffice/spreadsheet"
)

func Example_document() {
	// see the github.com/xmdas-link/unioffice/document documentation or _examples/document
	// for more examples
	doc := document.New()
	doc.AddParagraph().AddRun().AddText("Hello World!")
	doc.SaveToFile("document.docx")
}

func Example_spreadsheeet() {
	// see the github.com/xmdas-link/unioffice/spreadsheet documentation or _examples/spreadsheet
	// for more examples
	ss := spreadsheet.New()
	sheet := ss.AddSheet()
	sheet.AddRow().AddCell().SetString("Hello")
	sheet.Cell("B1").SetString("World!")
	ss.SaveToFile("workbook.xlsx")
}
