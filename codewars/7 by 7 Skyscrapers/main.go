package main

import (
	"fmt"
	"log"
)

const max int = 7

var bitMasks = []int{
	0x00,
	0x01,
	0x02,
	0x04,
	0x08,
	0x10,
	0x20,
	0x40,
	0x80,
}
var clues []int
var rows []int
var cols []int

func genPossiblePerms(clue int, predefined []int, curStep int, leftOvers []bool, curMax int, cur []int, perms *[][]int) {
	if curStep == max {
		var saveCur = make([]int, max)
		copy(saveCur, cur)
		*perms = append(*perms, saveCur)
		return
	}
	if predefined[curStep] > 0 {
		var element = predefined[curStep]
		var arrayElement = element - 1
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
		return
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
	if loc < max {
		x, y = loc_mod, i
	} else if loc < max*2 {
		x, y = max-1-i, loc_mod
	} else if loc < max*3 {
		x, y = max-1-loc_mod, max-1-i
	} else {
		x, y = i, max-1-loc_mod
	}
	return x, y
}

func solvePartial(clues []int, loc int, puzzle [][]int) bool {
	if loc == max*4 {
		return true
	}
	if clues[loc] == 0 && loc < max*4 {
		return solvePartial(clues, loc+1, puzzle)
	}
	fmt.Println(loc)
	printPuzzle(puzzle)

	var predefined = make([]int, max)
	leftOvers := []bool{true, true, true, true, true, true, true}

	var x int
	var y int
	for i := 0; i < max; i++ {
		x, y = getXY(loc, i)
		predefined[i] = puzzle[y][x]
		if predefined[i] > 0 {
			leftOvers[predefined[i]-1] = false
		}
	}
	cur := make([]int, max)
	perms := make([][]int, 0)
	genPossiblePerms(clues[loc], predefined, 0, leftOvers, 0, cur, &perms)

	for _, perm := range perms {
		valid := true
		for i := 0; i < max; i++ {
			x, y = getXY(loc, i)
			if predefined[i] == 0 && !unsetBits(x, y, perm[i]) {
				valid = false
				for j := 0; j < i; j++ {
					x, y = getXY(loc, j)
					setBits(x, y, perm[j])
				}
				break
			}
		}

		if valid {
			for i := 0; i < max; i++ {
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
	fmt.Println("<-- back from", loc)
	return false
}

func fillGaps(x int, y int, puzzle [][]int) bool {
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

	rows = []int{
		0x7F, 0x7F, 0x7F, 0x7F, 0x7F, 0x7F, 0x7F,
	}
	cols = []int{
		0x7F, 0x7F, 0x7F, 0x7F, 0x7F, 0x7F, 0x7F,
	}

	var puzzle [][]int
	for i := 0; i < max; i++ {
		var row = make([]int, max)
		puzzle = append(puzzle, row)
	}

	if !solvePartial(clues, 0, puzzle) {
		printPuzzle(puzzle)
		log.Fatal("Failed to solve puzzle")
	}
	fillGaps(0, 0, puzzle)
	printPuzzle(puzzle)
	return puzzle
}

func main() {
	// 3 [0 0 1 0] 0 [false true true true] 0 [0 0 0 0]
	// perms := make([][]int, 0)
	// var predefined = []int{0, 0, 3, 0}
	// var leftOvers = []bool{true, true, false, true}
	// var cur = []int{0, 0, 0, 0}
	// genPossiblePerms(3, predefined, 0, leftOvers, 0, cur, &perms)
	// fmt.Print(perms)
	SolvePuzzle(_clues)
}

func printPuzzle(puzzle [][]int) {
	fmt.Printf("    ")
	for i := 0; i < max; i++ {
		fmt.Printf("%-2v ", _clues[i])
	}
	fmt.Println()
	for i := 0; i < max; i++ {
		if i == 0 {
			fmt.Printf("%2v [", _clues[max*4-i-1])
		} else {
			fmt.Printf("%2v  ", _clues[max*4-i-1])
		}
		for j := 0; j < max; j++ {
			fmt.Printf("%-2v ", puzzle[i][j])
		}
		if i == max-1 {
			fmt.Printf("]%2v", _clues[i+max])
		} else {
			fmt.Printf(" %2v", _clues[i+max])
		}
		fmt.Printf("\n")
	}
	fmt.Print("    ")
	for i := 0; i < max; i++ {
		fmt.Printf("%-2v ", _clues[max*3-i-1])
	}
	fmt.Println()
}

var _clues = []int{
	7, 0, 0, 0, 2, 2, 3,
	0, 0, 3, 0, 0, 0, 0,
	3, 0, 3, 0, 0, 5, 0,
	0, 0, 0, 0, 5, 0, 4}

var result = [][]int{
	{1, 5, 6, 7, 4, 3, 2},
	{2, 7, 4, 5, 3, 1, 6},
	{3, 4, 5, 6, 7, 2, 1},
	{4, 6, 3, 1, 2, 7, 5},
	{5, 3, 1, 2, 6, 4, 7},
	{6, 2, 7, 3, 1, 5, 4},
	{7, 1, 2, 4, 5, 6, 3},
}
