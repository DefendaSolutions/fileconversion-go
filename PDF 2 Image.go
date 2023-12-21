/*
File Name:  PDF 2 Image.go
Copyright:  2019 Kleissner Investments s.r.o.
Author:     Peter Kleissner
*/

package fileconversion

import (
	"errors"
	"image"
	"io"
)

var xObjectImages = 0
var inlineImages = 0

// ImageResult contains an extracted image
type ImageResult struct {
	Image image.Image
	Name  string
}

// PDFExtractImages extracts all images from a PDF file
func PDFExtractImages(input io.ReadSeeker) (images []ImageResult, err error) {

	return nil, errors.New("Unhandled PDFExtractImages")
}

func extractImagesOnPage(page any) ([]any, error) {

	return nil, errors.New("Unhandled extractImagesOnPage")
}

func extractImagesInContentStream(contents string, resources *any) ([]any, error) {

	return nil, errors.New("Unhandled extractImagesInContentStream")
}
