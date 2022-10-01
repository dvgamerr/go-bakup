package main

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

type BakupGo struct {
	Archive     *os.File
	ZipFile     *zip.Writer
	RootInclude *bool
	Source      *string
	Destination *string
	Name        *string
}

func (bg *BakupGo) Close() {
	bg.ZipFile.Close()
	bg.Archive.Close()
}

func (bg *BakupGo) CreateZip(filename string) error {
	var err error
	Infof("Create Archive '%s'", filename)
	bg.Archive, err = os.Create(filepath.Join(*bg.Destination, filename))
	if err != nil {
		return err
	}
	bg.ZipFile = zip.NewWriter(bg.Archive)
	return nil
}

func (bg *BakupGo) AppendFile(file string) error {
	fileName := filepath.Base(file)
	if *bg.RootInclude {
		fileName = filepath.Join(filepath.Base(filepath.Dir(file)), filepath.Base(file))
	}
	Infof(" - Added '%s'", fileName)

	zipRead, err := os.Open(file)
	if err != nil {
		return err
	}
	defer zipRead.Close()

	zipWrite, err := bg.ZipFile.Create(fileName)
	if err != nil {
		return err
	}

	if _, err = io.Copy(zipWrite, zipRead); err != nil {
		return err
	}

	return nil
}
