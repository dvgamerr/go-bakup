package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	bakup := &BakupGo{}
	bakup.RootInclude = flag.Bool("r", false, "# target directory or file")
	bakup.Source = flag.String("src", "", "# target directory or file")
	bakup.Name = flag.String("name", "backup", "# backup filename")
	bakup.Destination = flag.String("dest", ".", "# target directory or file")
	flag.Parse()

	defer bakup.Close()

	if _, err := os.Stat(*bakup.Source); os.IsNotExist(err) {
		panic("source directory not exist")
	}

	if _, err := os.Stat(*bakup.Destination); os.IsNotExist(err) {
		panic("destination directory not exist")
	}

	zipName := fmt.Sprintf("%s_%s.zip", time.Now().Format("20060201"), *bakup.Name)
	if err := bakup.CreateZip(filepath.Join(*bakup.Destination, zipName)); err != nil {
		Error(err)
	}

	if err := filepath.Walk(*bakup.Source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		return bakup.AppendFile(path)
	}); err != nil {
		Error(err)
	}

	Info("Successful")
}
