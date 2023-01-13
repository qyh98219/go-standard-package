package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func TestScanner() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func ScannerTxt() {
	file, err := os.Open("./InputOutput/bufio-带缓存的IO/静夜思.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading file input:", err)
	}
}

func main() {
	ScannerTxt()
}
