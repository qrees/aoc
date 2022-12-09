package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var err error
	var buf []byte
	var count int
	var floor = 0

	buf = make([]byte, 1024)
	for err != io.EOF {
		count, err = os.Stdin.Read(buf)
		for _, char := range buf[:count] {
			switch char {
			case '(':
				floor += 1
			case ')':
				floor -= 1
			default:
				panic(fmt.Sprintf("Unexpected character in input: %v", char))
			}
		}
	}
	fmt.Printf("Answer: %v", floor)
}
