package main

import (
	"fmt"
	"math"
)

var powMap map[int][]int

// LastDigit ...
func LastDigit(as []int) int {
	fmt.Println("as", as)
	if len(as) == 0 {
		return 1
	}
	if len(as) == 1 {
		return int(math.Mod(float64(as[0]), float64(10)))
	}
	var lastIndex int
	var element int
	// var index int
	var base int
	powMap = make(map[int][]int)
	powMap[1] = []int{1}
	powMap[2] = []int{6, 2, 4, 8}
	powMap[3] = []int{1, 3, 9, 7}
	powMap[4] = []int{6, 4}
	powMap[5] = []int{5}
	powMap[6] = []int{6}
	powMap[7] = []int{1, 7, 9, 3}
	powMap[8] = []int{6, 8, 4, 2}
	powMap[9] = []int{1, 9}

	lastIndex = len(as) - 1
	element = as[lastIndex]
	for i := lastIndex; i > 1; i-- {
		base = as[i-1]
		if base == 0 {
			if element > 0 {
				element = 0
			} else {
				element = 1
			}
			continue
		}
		// baseCount = len(powMap[base])
		if element == 0 {
			element = 1
			continue
		}
		base = int(math.Mod(float64(base), 4))
		fmt.Print("Calc1: pow(", base, ", ", element, ") =")
		element = int(math.Pow(float64(base), float64(element)))
		element = int(math.Mod(float64(element), 4))
		fmt.Println(element)
	}
	return powMap[as[0]][int(math.Mod(float64(element), 4))]
}

func main() {
	// fmt.Println(LastDigit([]int{}))
	// fmt.Println(LastDigit([]int{0}))
	// fmt.Println(LastDigit([]int{0, 0}))
	// fmt.Println(LastDigit([]int{3, 4, 5}))
	fmt.Println(LastDigit([]int{7, 6, 21}))
	fmt.Println(LastDigit([]int{2, 3, 4, 7, 8}))
}
