package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var res uint
	var ribbon uint = 0
	var err error
	var reader *bufio.Reader
	var line []byte
	var dimensions []string
	var dim []int = []int{0, 0, 0}
	var input *os.File
	input, err = os.Open("AoC02.input")

	if err != nil {
		log.Fatal(err)
	}

	reader = bufio.NewReader(input)
	for true {
		line, _, err = reader.ReadLine()
		if err == io.EOF {
			break
		}
		dimensions = strings.Split(string(line), "x")
		for i, val := range dimensions {
			dim[i], err = strconv.Atoi(val)
		}

		sort.Ints(dim)
		var smallest int = dim[0] * dim[1]

		res += uint(2*dim[0]*dim[1] +
			2*dim[0]*dim[2] +
			2*dim[1]*dim[2] +
			smallest)
		ribbon += uint(dim[0]*2 + dim[1]*2 + dim[0]*dim[1]*dim[2])
		//os.Stdout.WriteString(fmt.Sprint(res));
	}
	os.Stdout.WriteString(fmt.Sprint(ribbon))
}
