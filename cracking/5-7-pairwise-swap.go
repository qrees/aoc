package main

import "fmt"

func swap_bits(value uint32) uint32 {
	var mask_even uint32 = 0xAAAAAAAA
	var mask_odd uint32 = 0x55555555

	return ((value & mask_even) >> 1) | ((value & mask_odd) << 1)
}

func main() {
	fmt.Println(swap_bits(1));
	fmt.Println(swap_bits(2));
	fmt.Println(swap_bits(10));
}
