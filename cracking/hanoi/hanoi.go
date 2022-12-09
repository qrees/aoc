package main

import "fmt"

func _paren(open int, closed int) {
	if open > 0 {
		fmt.Println("(")
		_paren(open-1, closed)
	}
	if open < closed {
		if closed > 0 {
			fmt.Println(")")
			_paren(open, closed-1)
		}
	}
	fmt.Println("-")
}

func paren(n int) {
	fmt.Println("(")
	_paren(n-1, n)
}

func main() {
	paren(2)
}
