package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func TestReader() {
	reader := strings.NewReader("Hello World")
	p := make([]byte, 1024)
	_, err := reader.Read(p)
	if err != nil {
		if err != io.EOF {
			log.Fatal(err)
			return
		}
	}
	fmt.Printf("data:%s", string(p))
}

func TestReaderAt() {
	reader := strings.NewReader("Go中文学习网")
	p := make([]byte, 100)
	n, err := reader.ReadAt(p, 2)
	if err != nil {
		if err != io.EOF {
			log.Fatal(err)
			return
		}
	}
	if n > 0 {
		fmt.Printf("data: %s \n", string(p))
	}
}

func TestWriteAt() {
	file, err := os.Create("writeAt.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("Golang中文社区--这里是多余")
	n, err := file.WriteAt([]byte("Go语言中文网"), 20)
	if err != nil {
		panic(err)
	}
	fmt.Printf("n: %v\n", n)
}

func TestReadFrom() {
	file, err := os.Open("writeAt.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(os.Stdout)
	writer.ReadFrom(file)
	writer.Flush()
}

func TestWriteTo() {
	reader := bytes.NewReader([]byte("Go语言中文网"))
	reader.WriteTo(os.Stdout)
}

func TestSeeker() {
	reader := strings.NewReader("Go语言中文网")
	_, err := reader.Seek(-6, io.SeekEnd)
	if err != nil {
		panic(err)
	}
	reader.WriteTo(os.Stdout)
}

func TestByteReaderByteWriter() {
	var ch byte
	fmt.Scanf("%c\n", &ch)

	buffer := new(bytes.Buffer)
	err := buffer.WriteByte(ch)
	if err == nil {
		fmt.Println("写入一个字节成功，准备读取...")
		newByte, err := buffer.ReadByte()
		if err != nil {
			fmt.Printf("读取一个字节失败, err: %s \n", err)
		}
		fmt.Printf("读取的字节: %c \n", newByte)
	} else {
		fmt.Println("写入错误")
	}
}

func TestLimitedReader() {
	content := "This is LimitReader example"
	reader := strings.NewReader(content)
	limitReader := &io.LimitedReader{R: reader, N: 8}
	for limitReader.N > 0 {
		tmp := make([]byte, 2)
		limitReader.Read(tmp)
		fmt.Printf("%s \n", tmp)
	}
}

func main() {
	TestLimitedReader()
}
