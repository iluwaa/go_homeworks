package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var pathToFile string
	var searchString string

	findStringInFile(pathToFile, searchString)
}

func findStringInFile(pathToFile string, searchString string) {
	fmt.Printf("Specify path to file: ")
	fmt.Scan(&pathToFile)

	isFileExist(pathToFile)

	fmt.Printf("\nType a string which ypu want to find in '%s': ", pathToFile)
	fmt.Scan(&searchString)

	content := readFile(pathToFile)
	for i, str := range strings.Split(string(content), "\n") {
		if strings.Contains(str, searchString) {
			fmt.Printf("Found string '%s', at line %d\n", searchString, i+1)
		}
	}
}

func readFile(pathToFile string) string {
	file, err := os.Open(pathToFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	return string(content)
}

func isFileExist(pathToFile string) {
	if _, err := os.Stat(pathToFile); os.IsNotExist(err) {
		fmt.Printf("File '%s' does not exist.\n", pathToFile)
		os.Exit(1)
	}
}
