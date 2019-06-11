package main

import (
	"fmt"
	"log"
)

const max int = 4

var bitMasks = []int{
	0x00,
	0x01,
	0x02,
	0x04,
	0x08,
}

var rows = []int{
	0xF, 0xF, 0xF, 0xF,
}

var cols = []int{
	0xF, 0xF, 0xF, 0xF,
}

func genPossiblePerms(clue int, predefined []int, curStep int, leftOvers []bool, curMax int, cur []int, perms *[][]int) {
	if curStep == max {
		var saveCur = []int{0, 0, 0, 0}
		copy(saveCur, cur)
		*perms = append(*perms, saveCur)
		return
	}
	if predefined[curStep] > 0 {
		var element = predefined[curStep]
		var arrayElement = element - 1
		if !leftOvers[arrayElement] {
			return
		}
		if element > max-clue+1 {
			return
		}
		leftOvers[arrayElement] = false
		cur[curStep] = element
		if element > curMax && clue > 0 {
			genPossiblePerms(clue-1, predefined, curStep+1, leftOvers, element, cur, perms)
		}
		if element < curMax {
			genPossiblePerms(clue, predefined, curStep+1, leftOvers, curMax, cur, perms)
		}
		if element == curMax {
			log.Fatal("Duplicate value")
		}
		leftOvers[arrayElement] = true
	}
	for arrayElement, exists := range leftOvers {
		var element = arrayElement + 1
		if !exists {
			continue
		}
		if element > max-clue+1 {
			continue
		}
		leftOvers[arrayElement] = false
		cur[curStep] = element
		if element > curMax && clue > 0 {
			genPossiblePerms(clue-1, predefined, curStep+1, leftOvers, element, cur, perms)
		}
		if element < curMax {
			genPossiblePerms(clue, predefined, curStep+1, leftOvers, curMax, cur, perms)
		}
		if element == curMax {
			log.Fatal("Duplicate value")
		}
		leftOvers[arrayElement] = true
	}
}

// func verify(puzzle [][]int) bool {
// 	for i := 0; i < max; i++ {
// 		var bitmask int
// 		for j := 0; j < max; j++ {
// 			if puzzle[i][j] > 0 {
// 				if bitmask&(bitMasks[puzzle[i][j]]) == 0 {
// 					bitmask = bitmask | bitMasks[puzzle[i][j]]
// 				} else {
// 					return false
// 				}
// 			}
// 		}
// 	}
// 	for i := 0; i < max; i++ {
// 		var bitmask int
// 		for j := 0; j < max; j++ {
// 			if puzzle[j][i] > 0 {
// 				if bitmask&(bitMasks[puzzle[j][i]]) == 0 {
// 					bitmask = bitmask | bitMasks[puzzle[j][i]]
// 				} else {
// 					return false
// 				}
// 			}
// 		}
// 	}
// 	return true
// }

func printPuzzle(puzzle [][]int) {
	fmt.Printf("    ")
	for i := 0; i < max; i++ {
		fmt.Printf("%-2v ", cols[i])
	}
	fmt.Println()
	for i := 0; i < max; i++ {
		if i == 0 {
			fmt.Printf("%2v [", rows[i])
		} else {
			fmt.Printf("%2v  ", rows[i])
		}
		for j := 0; j < max; j++ {
			fmt.Printf("%-2v ", puzzle[i][j])
		}
		if i == max-1 {
			fmt.Printf("]")
		}
		fmt.Printf("\n")
	}
}

func unsetBits(x int, y int, value int) bool {
	if cols[x]&bitMasks[value] == 0 {
		return false
	}
	if rows[y]&bitMasks[value] == 0 {
		return false
	}
	cols[x] = cols[x] &^ bitMasks[value]
	rows[y] = rows[y] &^ bitMasks[value]
	return true
}

func setBits(x int, y int, value int) {
	cols[x] = cols[x] | bitMasks[value]
	rows[y] = rows[y] | bitMasks[value]
}

func getXY(loc int, i int) (int, int) {
	var x int
	var y int
	var loc_mod = loc % max
	if loc < 4 {
		x, y = loc_mod, i
	} else if loc < 8 {
		x, y = 3-i, loc_mod
	} else if loc < 12 {
		x, y = 3-loc_mod, 3-i
	} else {
		x, y = i, 3-loc_mod
	}
	return x, y
}

// func checkBits(puzzle [][]int) {
// 	var _cols = []int{0xF, 0xF, 0xF, 0xF}
// 	var _rows = []int{0xF, 0xF, 0xF, 0xF}
// 	for x := 0; x < max; x++ {
// 		for y := 0; y < max; y++ {
// 			value := puzzle[y][x]
// 			if value != 0 {
// 				_cols[x] = _cols[x] &^ bitMasks[value]
// 				_rows[y] = _rows[y] &^ bitMasks[value]
// 			}
// 		}
// 	}
// 	for i := 0; i < max; i++ {
// 		if _cols[i] != cols[i] {
// 			printPuzzle(puzzle)
// 			log.Fatal("error")
// 		}
// 		if _rows[i] != rows[i] {
// 			printPuzzle(puzzle)
// 			log.Fatal("error")
// 		}
// 	}
// }

func solvePartial(clues *[]int, loc int, puzzle [][]int) bool {
	// if !verify(puzzle) {
	// 	return false
	// }
	if loc == 16 {
		return true
	}
	if (*clues)[loc] == 0 && loc < 16 {
		return solvePartial(clues, loc+1, puzzle)
	}

	var predefined = []int{0, 0, 0, 0}
	leftOvers := []bool{true, true, true, true}

	var x int
	var y int
	for i := 0; i < 4; i++ {
		x, y = getXY(loc, i)
		predefined[i] = puzzle[y][x]
	}
	cur := []int{0, 0, 0, 0}
	perms := make([][]int, 0)
	genPossiblePerms((*clues)[loc], predefined, 0, leftOvers, 0, cur, &perms)

	for _, perm := range perms {
		valid := true
		for i := 0; i < 4; i++ {
			x, y = getXY(loc, i)
			if predefined[i] == 0 && !unsetBits(x, y, perm[i]) {
				valid = false
				for j := 0; j < i; j++ {
					x, y = getXY(loc, j)
					setBits(x, y, perm[j])
				}
				// fmt.Printf("Cannot use permutation %v at %v in puzzle because %v %v %v:\n", perm, loc, x, y, i)
				// printPuzzle(puzzle)
				break
			}
		}

		if valid {
			// fmt.Printf("--> Using permutation %v at %v in puzzle:\n", perm, loc)
			// printPuzzle(puzzle)
			for i := 0; i < 4; i++ {
				x, y = getXY(loc, i)
				puzzle[y][x] = perm[i]
			}
			if solvePartial(clues, loc+1, puzzle) {
				return true
			}
			for i := 0; i < max; i++ {
				x, y = getXY(loc, i)
				if predefined[i] == 0 {
					setBits(x, y, perm[i])
				}
			}
		}
	}
	for i := 0; i < max; i++ {
		x, y = getXY(loc, i)
		puzzle[y][x] = predefined[i]
	}

	return false
}

func fillGaps(x int, y int, puzzle [][]int) bool {
	// fmt.Printf("%v %v\n", x, y)
	if y == max {
		return true
	}
	if puzzle[y][x] == 0 {
		var possible = cols[x] & rows[y]
		for i := 0; i < max; i++ {
			value := i + 1
			if possible&(1<<uint(i)) != 0 {
				puzzle[y][x] = value
				unsetBits(x, y, value)
				res := fillGaps((x+1)%max, y+((x+1)/max), puzzle)
				if res {
					return true
				} else {
					setBits(x, y, value)
					puzzle[y][x] = 0
				}
			}
		}
		return false
	}
	return fillGaps((x+1)%max, y+((x+1)/max), puzzle)
}

func SolvePuzzle(clues []int) [][]int {
	var puzzle = [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	solvePartial(&clues, 0, puzzle)
	fillGaps(0, 0, puzzle)
	return puzzle
}

func main() {
	puzzle := SolvePuzzle(clues)
	printPuzzle(puzzle)
}

var clues = []int{
	2, 2, 1, 3,
	2, 2, 3, 1,
	1, 2, 2, 3,
	3, 2, 1, 3}

// var clues = []int{
// 	0, 0, 1, 2,
// 	0, 2, 0, 0,
// 	0, 3, 0, 0,
// 	0, 1, 0, 0}

var outcome = [][]int{
	{2, 1, 4, 3},
	{3, 4, 1, 2},
	{4, 2, 3, 1},
	{1, 3, 2, 4}}
