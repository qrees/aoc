package main

import (
    "os"
    "bufio"
    "strconv"
    "fmt"
)


func main(){
    var input *os.File;
    var scanner *bufio.Scanner;
    var line string;
    //var unquoted string;
    var quoted string;
    //var err error;
    var mem_count int;

    input = os.Stdin;
    scanner = bufio.NewScanner(input);

    for scanner.Scan() {
        line = scanner.Text();
        fmt.Println(line);
        //unquoted, err = strconv.Unquote(line);
        //if err != nil {
        //    fmt.Println("Failed to unquote", err);
        //    panic(line);
        //}
        quoted = strconv.Quote(line);
        mem_count += len(quoted);
        mem_count -= len(line);
    }
    println(mem_count);
}
