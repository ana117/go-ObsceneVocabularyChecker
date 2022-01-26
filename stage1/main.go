package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// reading file name input
	var fileName string
	fmt.Scan(&fileName)

	// opening file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// create scanner and split by words
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	// printing each words
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
