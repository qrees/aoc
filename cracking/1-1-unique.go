package main

import (
    "fmt"
    "io"
    "os"
)

func main(){
    var buf []byte;
    var err error;
    var count int;
    var counts map[byte]bool = make(map[byte]bool);
    var unique bool;
    //var char byte;

    unique = true;
	buf = make([]byte, 1024)
	for err != io.EOF {
		count, err = os.Stdin.Read(buf)
		for _, char := range buf[:count] {
            if (counts[char]) {
                unique = false;
                break;
            }
            counts[char] = true;
		}
	}

    if (unique) {
        fmt.Println("Is unique");
    } else {
        fmt.Println("Is NOT unique");
    }
}
