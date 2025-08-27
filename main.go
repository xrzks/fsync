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

	exists, err := validateDirectory(sourceDirectory)
	if err != nil {
		return fmt.Errorf("failed to validate source directory %q: %w", sourceDirectory, err)
	}
	if !exists {
		return fmt.Errorf("source directory %q does not exist", sourceDirectory)
	}
	exists, err = validateDirectory(targetDirectory)
	if err != nil {
		return fmt.Errorf("failed to validate target directory %q: %w", targetDirectory, err)
	}
	if !exists {
		return fmt.Errorf("target directory %q does not exist", targetDirectory)
	}

	sourceEntries, err := os.ReadDir(sourceDirectory)
	if err != nil {
		return fmt.Errorf("failed to read source directory %q: %w", sourceDirectory, err)
	}

	for _, sourceEntry := range sourceEntries {
		sourceFilePath := sourceDirectory + sourceEntry.Name()
		targetFilePath := targetDirectory + sourceEntry.Name()

		err = copyFile(sourceFilePath, targetFilePath)
		if err != nil {
			return fmt.Errorf("failed to copy file %q to %q: %w", sourceFilePath, targetFilePath, err)
		}
	}

	return nil
}

func validateDirectory(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	if info.IsDir() {
		return true, nil
	} else {
		return false, nil
	}
}

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
