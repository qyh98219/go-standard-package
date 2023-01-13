package main

import (
	"bufio"
	"fmt"
	"strings"
)

func TestReadSlice() {
	reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \n It is the home of gophers"))
	line, _ := reader.ReadBytes('\n')
	fmt.Printf("the line:%s\n", line)
	n, _ := reader.ReadBytes('\n')
	fmt.Printf("the line:%s\n", line)
	fmt.Println(string(n))
}

func main() {
	TestReadSlice()
}
