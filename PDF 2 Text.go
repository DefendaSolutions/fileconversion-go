/*
File Name:  PDF 2 Text.go
Copyright:  2018 Kleissner Investments s.r.o.
Author:     Peter Kleissner

This code uses the commercial library UniDoc https://unidoc.io/ to extract text from PDFs.
*/

package fileconversion

import (
	"errors"
	"io"
	"time"
)

// PDFListContentStreams writes all text streams in a PDF to the writer
// It returns the number of characters attempted written (excluding "Page N" and new-lines) and an error, if any. It can be used to determine whether any text was extracted.
// The parameter size is the max amount of bytes (not characters) to write out.
func PDFListContentStreams(f io.ReadSeeker, w io.Writer, size int64) (written int64, err error) {

	return 0, errors.New("Unhandled PDF 2 Text")

}

// PDFGetCreationDate tries to get the creation date
func PDFGetCreationDate(f io.ReadSeeker) (date time.Time, valid bool) {

	return date, false

}
