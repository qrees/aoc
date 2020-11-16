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

func main() {
    var arr []uint
    var input *os.File
    var err error
    var curXLoc, maxXLoc, minXLoc, curYLoc, maxYLoc, minYLoc int64
    var count int64;
    var scanner *bufio.Scanner

    input, err = os.Open("AoC03.input")
    if err != nil {
        log.Fatal(err)
    }
    scanner = bufio.NewScanner(input)

    split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
        return 1, data[0:1], nil
    }

    scanner.Split(split)

    for scanner.Scan() {
        var direction []byte = scanner.Bytes()
        if direction[0] == '>' {
            curXLoc += 1
            maxXLoc = max(maxXLoc, curXLoc)
        }
        if direction[0] == '<' {
            curXLoc -= 1
            minXLoc = min(minXLoc, curXLoc)
        }
        if direction[0] == 'v' {
            curYLoc += 1
            maxYLoc = max(maxYLoc, curYLoc)
        }
        if direction[0] == '^' {
            curYLoc -= 1
            minYLoc = min(minYLoc, curYLoc)
        }
    }
    arr = make([]uint, (maxYLoc-minYLoc+1)*(maxXLoc-minXLoc+1))

    var row int64 = (maxXLoc - minXLoc + 1)
    fmt.Println(row)
    minXLoc = -minXLoc
    minYLoc = -minYLoc
    arr[minXLoc+(row*minYLoc)] = 1
    input, err = os.Open("AoC03.input")
    scanner = bufio.NewScanner(input)
    scanner.Split(split)

    curXLoc = 0;
    curYLoc = 0;
    count = 1;
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

        if (arr[minXLoc+curXLoc+(row*(minYLoc+curYLoc))] == 0) {
            count += 1;
        }
        // fmt.Println("cur: ", minXLoc+curXLoc+(row*(minYLoc+curXLoc)))
        // fmt.Println("X: ", minXLoc+curXLoc)
        arr[minXLoc+curXLoc+(row*(minYLoc+curYLoc))]++
    }
    fmt.Println();
    fmt.Println(count)
    os.Exit(0)
}
