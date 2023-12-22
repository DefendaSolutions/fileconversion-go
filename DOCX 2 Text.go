/*
File Name:  DOCX 2 Text.go
Copyright:  2018 Kleissner Investments s.r.o.
Author:     Peter Kleissner

This code is forked from https://github.com/guylaor/goword and extracts text from DOCX files.
*/

package fileconversion

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// models.go

// WordDocument is a full word doc
type WordDocument struct {
	Paragraphs []WordParagraph `xml:"p"`
}

// WordParagraph is a single paragraph
type WordParagraph struct {
	Style      WordStyle       `xml:"pPr>pStyle"`
	Rows       []WordRow       `xml:"r"`
	Hyperlinks []WordHyperlink `xml:"hyperlink"`
}

// WordStyle ...
type WordStyle struct {
	Val string `xml:"val,attr"`
}

// WordRow ...
type WordRow struct {
	Text string `xml:"t"`
}

// WordHyperlink ...
type WordHyperlink struct {
	Id   string    `xml:"id,attr"`
	Rows []WordRow `xml:"r"`
}

// AsText returns all text in the document
func (w WordDocument) AsText() string {
	text := ""
	for _, v := range w.Paragraphs {
		if len(v.Rows) > 0 {
			for _, rv := range v.Rows {
				text += rv.Text
			}
			text += "\n"
		}
		if len(v.Hyperlinks) > 0 {
			for _, hl := range v.Hyperlinks {
				for _, rv := range hl.Rows {
					text += rv.Text
				}
			}
			text += "\n"
		}
	}
	return text
}

// goword.go

// DOCX2Text extracts text of a Word document
// Size is the full size of the input file.
func DOCX2Text(file io.ReaderAt, size int64) (string, error) {

	doc, err := openWordFile(file, size)
	if err != nil {
		return "", err
	}

	docx, err := WordParse(doc)
	if err != nil {
		return "", err
	}

	return docx.AsText(), nil
}

// WordParse parses a word file
func WordParse(doc string) (WordDocument, error) {

	docx := WordDocument{}
	r := strings.NewReader(string(doc))
	decoder := xml.NewDecoder(r)

	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "p" {
				var p WordParagraph
				decoder.DecodeElement(&p, &se)
				docx.Paragraphs = append(docx.Paragraphs, p)
			}
		}
	}
	return docx, nil
}

func openWordFile(file io.ReaderAt, size int64) (string, error) {

	// Open a zip archive for reading. word files are zip archives
	r, err := zip.NewReader(file, size)
	if err != nil {
		return "", err
	}

	retvalue := ""

	// Iterate through the files in the archive,
	// find document.xml
	for _, f := range r.File {

		//fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			return "", err
		}
		defer rc.Close()
		if f.Name == "word/document.xml" {
			doc, err := io.ReadAll(rc)
			if err != nil {
				return "", err
			}
			retvalue += fmt.Sprintf("%s", doc)
		}

		// ! others .docx in the word/ folder should be parsed as ZIP to handle nested Word documents
		if strings.ToLower(filepath.Ext(f.Name)) == ".docx" {
			fmt.Fprintf(os.Stderr, "Unhandled nested %s in .docx file\n", f.Name)
		}
	}

	return retvalue, nil
}

// IsFileDOCX checks if the data indicates a DOCX file
// DOCX has a signature of 50 4B 03 04
func IsFileDOCX(data []byte) bool {
	return bytes.HasPrefix(data, []byte{0x50, 0x4B, 0x03, 0x04})
}
