package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	// reading taboo word input
	var tabooWord string
	fmt.Scan(&tabooWord)
	tabooWord = strings.ToLower(tabooWord)

	// check each words
	var tabooWordExist bool
	for scanner.Scan() {
		if strings.ToLower(scanner.Text()) == tabooWord {
			tabooWordExist = true
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(tabooWordExist)
}
