package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
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

	// putting forbidden words to set (map)
	tabooMap := make(map[string]struct{})
	for scanner.Scan() {
		// add to set if it doesn't exist
		tabooWord := strings.ToLower(scanner.Text())
		if _, ok := tabooMap[tabooWord]; !ok {
			tabooMap[tabooWord] = struct{}{}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	var word string
	fmt.Scan(&word)
	for word != "exit" {
		// if word is a tabooWord, censor it
		if _, ok := tabooMap[strings.ToLower(word)]; ok {
			word = strings.Repeat("*", utf8.RuneCountInString(word))
		}
		fmt.Println(word)

		fmt.Scan(&word)
	}

	fmt.Println("Bye!")
}
