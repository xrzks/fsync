package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	sourceDirectory := "/Users/simon/dev/private/fsync/testSrcDir/"
	targetDirectory := "/Users/simon/dev/private/fsync/testTargetDir/"

	sourceEntries, err := os.ReadDir(sourceDirectory)
	if err != nil {
		return fmt.Errorf("failed to read source directory %q: %w", sourceDirectory, err)
	}

	for index, sourceEntry := range sourceEntries {
		sourceFilePath := sourceDirectory + sourceEntry.Name()
		targetFilePath := targetDirectory + sourceEntry.Name()

		fmt.Printf("%d 󰉋  %s\n", index, sourceFilePath)

		err = copyFile(sourceFilePath, targetFilePath)
		if err != nil {
			return fmt.Errorf("failed to copy file %q to %q: %w", sourceFilePath, targetFilePath, err)
		}
	}

	return nil
}

func copyFile(srcFile, targetFile string) error {
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
