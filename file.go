package main

import (
	"io"
	"os"
)

func copyFile(srcFile string, targetFile string) error {
	// open source file
	sourceFile, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// create target file
	destFile, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// copy contents
	_, err = io.Copy(destFile, sourceFile)
	return err
}
