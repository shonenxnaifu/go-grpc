package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(strings.NewReader("Hello, bufio packageee!\nHow are you doing?"))

	// Reading byte slice
	data := make([]byte, 20)
	n, err := reader.Read(data)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	fmt.Printf("Read %d bytes: %s\n", n, data[:n])

	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading string:", err)
	}
	fmt.Println("Read string:", line)

	writer := bufio.NewWriter(os.Stdout)

	// Writing byte slice
	data1 := []byte("Hello, bufio package!\n")
	n1, err := writer.Write(data1)
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", n1)

	// FLush the buffer to ensure all data is written to os.Stdout
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing write:", err)
		return
	}

	// Writing string
	str := "This is a string.\n"
	n2, err := writer.WriteString(str)
	if err != nil {
		fmt.Println("Error writing string:", err)
		return
	}
	fmt.Printf("Wrote %d bytes.\n", n2)

	// Flush the buffer
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}
}
