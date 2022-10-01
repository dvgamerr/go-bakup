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

	flag.StringVar(&bakup.Name, "name", "backup", "# backup filename")
	flag.BoolVar(&bakup.RootInclude, "r", false, "# target directory or file")
	flag.StringVar(&bakup.Source, "src", "", "# target directory or file")
	flag.StringVar(&bakup.Destination, "dest", ".", "# target directory or file")
	flag.Parse()

	defer bakup.Close()

	if _, err := os.Stat(bakup.Source); os.IsNotExist(err) {
		if os.Args[1] != "" {
			if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
				panic("source directory not exist")
			}
			bakup.Source = os.Args[1]
		}
	}

	if _, err := os.Stat(bakup.Destination); os.IsNotExist(err) {
		panic("destination directory not exist")
	}

	zipName := fmt.Sprintf("%s_%s.zip", time.Now().Format("20060201"), bakup.Name)
	if err := bakup.CreateZip(filepath.Join(bakup.Destination, zipName)); err != nil {
		Error(err)
	}

	if err := filepath.Walk(bakup.Source, func(path string, info os.FileInfo, err error) error {
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
