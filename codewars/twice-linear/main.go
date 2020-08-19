package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func DblLinear(n int) int {
	u := make([]int, 0)
	ai := 0
	bi := 0
	u = append(u, 1)

	for min(ai, bi) < n {
		x := 2*u[ai] + 1
		y := 3*u[bi] + 1
		if x > y {
			u = append(u, y)
			bi++
		}
		if y > x {
			u = append(u, x)
			ai++
		}
		if y == x {
			u = append(u, x)
			ai++
			bi++
		}
	}
	return u[n]
}

func main() {
	fmt.Print(DblLinear(10))
}
