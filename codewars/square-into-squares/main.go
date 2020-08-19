package main

import (
	"fmt"
	"math"
)

var res []int64

func min(a int64, b int64) int64 {
	if a > b {
		return b
	}
	return a
}

func DecomposeMore(n int64, previous int64) bool {
	if previous == 1 {
		return false
	}
	if n < 0 {
		return false
	}
	if n == 0 {
		return true
	}
	if n == 1 {
		fmt.Println("-> 1")
		res = append(res, 1)
		return true
	}
	sqrt := int64(math.Floor(math.Sqrt(float64(n))))
	cur := min(sqrt, previous-1)
	for ; cur > 0; cur-- {
		leftover := n - cur*cur
		if DecomposeMore(leftover, cur) {
			res = append(res, cur)
			fmt.Println("-> ", cur)
			return true
		}
	}
	return false
}

func Decompose(n int64) []int64 {
	res = make([]int64, 0)
	value := n * n
	sqrt := int64(math.Floor(math.Sqrt(float64(value - 1))))
	for ; sqrt > 0; sqrt-- {
		fmt.Println("->", sqrt)
		leftover := value - sqrt*sqrt
		if DecomposeMore(leftover, sqrt) {
			res = append(res, sqrt)
			return res
		}
	}
	return nil
}

func main() {
	Decompose(50)
}
