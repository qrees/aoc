package main

import (
	"fmt"
	"strconv"
)

func MaxRot(n int64) int64 {
	var s string = strconv.FormatInt(n, 10)
	var array []byte = []byte(s)
	var length int = len(array)
	var loc int = 0
	var max int64 = n

	for loc < length {
		var c byte = array[loc]
		var i int = loc
		for ; i < length-1; i++ {
			array[i] = array[i+1]
		}
		array[i] = c
		result := string(array)
		first, _ := strconv.ParseInt(result, 10, 64)
		if first > max {
			max = first
		}
		loc++
	}
	return max
}

func main() {
	fmt.Printf("%v\n", MaxRot(33331455060013))
	fmt.Printf("33456033503101")
}
