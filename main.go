package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/scanner"
)

func tokenize(source string) []string {
	var s scanner.Scanner
	s.Init(strings.NewReader(source))
	var tokens []string
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		tokens = append(tokens, s.TokenText())
	}
	return tokens
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run VMTranslator.go <filename>")
		os.Exit(1)
	}
	filename := os.Args[1]
	inputFilename, err := filepath.Abs(filename)
	if err != nil {
		fmt.Printf("Error resolving input file path: %v\n", err)
		os.Exit(1)
	}

	file, err := os.Open(inputFilename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	outputFile, err := os.Create("/home/kali/Desktop/output.asm")
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		_, err := writer.WriteString("dachi wrote " + line + "\n")
		fmt.Println(strings.Join(tokenize(line), " "))
		if err != nil {
			fmt.Printf("Error writing to output file: %v\n", err)
			os.Exit(1)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("File copied successfully")
}
