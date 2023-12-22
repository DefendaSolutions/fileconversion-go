package xls

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"
)

func TestIssue47(t *testing.T) {
	testdatapath := "testdata"

	dir, err := os.Open(testdatapath)
	if err != nil {
		t.Fatalf("Cant open testdata directory: %s", err)
	}

	files, err := dir.ReadDir(0)
	if err != nil {
		t.Fatalf("Cant read testdata directory contents: %s", err)
	}
	for _, f := range files {
		if filepath.Ext(f.Name()) == ".xls" {
			xlsfilename := f.Name()
			xlsxfilename := strings.TrimSuffix(xlsfilename, filepath.Ext(xlsfilename)) + ".xlsx"
			err := CompareXlsXlsx(path.Join(testdatapath, xlsfilename),
				path.Join(testdatapath, xlsxfilename))
			if err != "" {
				t.Fatalf("XLS file %s an XLSX file are not equal: %s", xlsfilename, err)
			}

		}
	}

}
