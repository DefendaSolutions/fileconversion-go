/*
File Name:  PDF 2 Text.go
Copyright:  2018 Kleissner Investments s.r.o.
Author:     Peter Kleissner

This code uses the commercial library UniDoc https://unidoc.io/ to extract text from PDFs.
*/

package fileconversion

import (
	"bytes"
	"io"
	"time"

	"github.com/ledongthuc/pdf"
)

// PDF2Text returns the plaintext content of the PDF file
// The parameter size is the PDF file size
func PDF2Text(f io.ReaderAt, size int64) (string, error) {

	r, err := pdf.NewReader(f, size)
	if err != nil {
		return "", err
	}

	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}

	buff := bytes.NewBuffer([]byte{})
	_, err = buff.ReadFrom(b)

	return buff.String(), err
}

// PDFGetCreationDate tries to get the creation date
func PDFGetCreationDate(f io.ReadSeeker) (date time.Time, valid bool) {

	return date, false

}
