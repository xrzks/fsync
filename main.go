package main

import (
	"fmt"
	"os"
)

func main() {
	sourceDirectory := "/Users/simon/dev/private/fsync/testSrcDir/"
	targetDirectory := "/Users/simon/dev/private/fsync/testTargetDir/"

	err := fileSync(sourceDirectory, targetDirectory)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
