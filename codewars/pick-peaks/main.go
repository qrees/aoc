package main

import (
	"fmt"
)

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(array []int) PosPeaks {
	if len(array) == 0 {
		return PosPeaks{
			Pos:   []int{},
			Peaks: []int{}}
	}
	var start_peak = -1
	var end_peak = -1
	var last_value = array[0]
	var pos = []int{}
	var peaks = []int{}

	for i := 1; i < len(array); i++ {
		var cur_value = array[i]
		if cur_value > last_value {
			start_peak = i
			end_peak = -1
		}
		if cur_value < last_value {
			end_peak = i - 1
		}
		if start_peak > -1 && end_peak > -1 {
			pos = append(pos, start_peak)
			peaks = append(peaks, last_value)
			start_peak = -1
			end_peak = -1
		}
		last_value = cur_value
	}
	return PosPeaks{
		Pos:   pos,
		Peaks: peaks}
}

func main() {
	fmt.Print(
		PickPeaks(
			[]int{2, 1, 3, 1, 2, 2, 2, 2, 1}))
	// PosPeaks{Pos: []int{2, 4}, Peaks: []int{3, 2}},
}
