package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func max(a int64, b int64) (max int64) {
    if a > b {
        return a
    } else {
        return b
    }
}

func min(a int64, b int64) (max int64) {
    if a < b {
        return a
    } else {
        return b
    }
}

func getminmax(scanner *bufio.Scanner) (int64, int64, int64, int64) {
    var curXLoc, maxXLoc, minXLoc, curYLoc, maxYLoc, minYLoc int64

    for scanner.Scan() {
        var direction []byte = scanner.Bytes()
        if direction[0] == '>' {
            curXLoc += 1
            maxXLoc = max(maxXLoc, curXLoc)
        } else if direction[0] == '<' {
            curXLoc -= 1
            minXLoc = min(minXLoc, curXLoc)
        } else if direction[0] == 'v' {
            curYLoc += 1
            maxYLoc = max(maxYLoc, curYLoc)
        } else if direction[0] == '^' {
            curYLoc -= 1
            minYLoc = min(minYLoc, curYLoc)
        } else {
            panic("Shit happened");
        }
    }
    return maxXLoc, minXLoc, maxYLoc, minYLoc
}

var arr []uint
var row int64
var count int64 = 1

func scan(scanner *bufio.Scanner, minXLoc int64, minYLoc int64){
    var curXLoc, curYLoc int64
    for scanner.Scan() {
        var direction []byte = scanner.Bytes()
        if direction[0] == '>' {
            curXLoc += 1
        }
        if direction[0] == '<' {
            curXLoc -= 1
        }
        if direction[0] == 'v' {
            curYLoc += 1
        }
        if direction[0] == '^' {
            curYLoc -= 1
        }

        fmt.Println(minXLoc+curXLoc, minYLoc+curYLoc);
        if (arr[minXLoc+curXLoc+(row*(minYLoc+curYLoc))] == 0) {
            count += 1;
        }
        arr[minXLoc+curXLoc+(row*(minYLoc+curYLoc))]++
    }

}

func main() {
    var input *os.File
    var maxXLoc, minXLoc, maxYLoc, minYLoc int64
    var maxXLoc2, minXLoc2, maxYLoc2, minYLoc2 int64
    var err error
    var scanner *bufio.Scanner

    input, err = os.Open("AoC03.input")
    if err != nil {
        log.Fatal(err)
    }

    scanner = bufio.NewScanner(input)
    split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
        return 2, data[0:1], nil
    }

    scanner.Split(split)
    maxXLoc, minXLoc, maxYLoc, minYLoc = getminmax(scanner);
    fmt.Println("First getminmax");
    input.Seek(1, 0)
    scanner = bufio.NewScanner(input)
    scanner.Split(split)
    maxXLoc2, minXLoc2, maxYLoc2, minYLoc2 = getminmax(scanner);
    fmt.Println("second getminmax");

    minXLoc = min(minXLoc, minXLoc2)
    maxXLoc = max(maxXLoc, maxXLoc2)
    minYLoc = min(minYLoc, minYLoc2)
    maxYLoc = max(maxYLoc, maxYLoc2)
    fmt.Println("Size:", (maxYLoc-minYLoc+1)*(maxXLoc-minXLoc+1));
    arr = make([]uint, (maxYLoc-minYLoc+1)*(maxXLoc-minXLoc+1))
    fmt.Println(minXLoc, maxXLoc, minYLoc, maxYLoc)

    row = (maxXLoc - minXLoc + 1)
    minXLoc = -minXLoc
    minYLoc = -minYLoc
    arr[minXLoc+(row*minYLoc)] = 1
    input.Seek(0, 0)
    scanner = bufio.NewScanner(input)
    scanner.Split(split)

    scan(scanner, minXLoc, minYLoc)

    input.Seek(1, 0)
    scanner = bufio.NewScanner(input)
    scanner.Split(split)
    scan(scanner, minXLoc, minYLoc)
    fmt.Println();
    fmt.Print(count)
    os.Exit(0)
}
