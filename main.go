package main

import (
	"fmt"
	"os"
)

func main() {
	err := readFromFile()
	if err != nil {
		os.Exit(1)
	}
}

func readFromFile() error {
	file, err := os.Create("one.txt")
	if err != nil {
		return fmt.Errorf("error reading file %w", err)
	}
	defer file.Close()
	return nil
}
