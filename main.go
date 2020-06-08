/*
 * Copyright (c) 2020. Author by KTECH Inc team. khang_duy.p
 */

package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.OpenFile("ziptesttt.zip", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		panic(err)
	}
	w := zip.NewWriter(f)

	err = appendFiles("ziptest/aa.csv", w)
	if err != nil {
		panic(err)
	}
	err = appendFiles("ziptest/he.json", w)
	if err != nil {
		panic(err)
	}
	defer w.Close()

}

func appendFiles(filename string, zipw *zip.Writer) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Failed to open %s: %s", filename, err)
	}
	defer file.Close()

	wr, err := zipw.Create(filename)
	if err != nil {
		msg := "Failed to create entry for %s in zip file: %s"
		return fmt.Errorf(msg, filename, err)
	}

	if _, err := io.Copy(wr, file); err != nil {
		return fmt.Errorf("Failed to write %s to zip: %s", filename, err)
	}

	return nil
}
