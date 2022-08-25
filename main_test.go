package recipebot

import (
	"archive/zip"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestMain_zip(t *testing.T) {
	if err := updateZip(); err != nil {
		t.Fatal(err)
	}

	zr, err := zip.OpenReader("../recipebot.zip")
	if err != nil {
		t.Fatal(err)
	}

	defer zr.Close()
}

func updateZip() error {
	w, err := os.OpenFile("../recipebot.zip", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	defer w.Close()
	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	return filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			switch info.Name() {
			case ".git":
				return filepath.SkipDir
			}
		} else {
			switch info.Name() {
			case ".DS_Store", ".gitignore", "test.env":
			default:
				if !strings.HasSuffix(info.Name(), "_test.go") {
					if err := addFileToZip(zipWriter, path); err != nil {
						return err
					}
				}
			}
		}

		return nil
	})
}

func addFileToZip(zw *zip.Writer, path string) error {
	fi, err := os.Lstat(path)
	if err != nil {
		return err
	}
	fh, err := zip.FileInfoHeader(fi)
	if err != nil {
		return err
	}
	fh.Method = zip.Deflate
	fh.Name = path
	w, err := zw.CreateHeader(fh)
	if err != nil {
		return err
	}
	r, err := os.Open(path)
	if err != nil {
		return err
	}
	defer r.Close()

	_, err = io.Copy(w, r)
	return err
}
