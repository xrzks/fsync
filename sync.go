package main

import (
	"fmt"
	"os"
)

func fileSync(sourceDirectory string, targetDirectory string) error {
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
