package main

import (
	"bytes"
	"fmt"
	"math"
)

// You may get original char by morse code like this: MORSE_CODE[char]

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func DecodeBits(bits string) string {
	skipping_zeros := true
	prev := '1'
	seq_start := 0
	seq_len := 0
	min_ := math.MaxInt64
	max_ := 0
	pos := 0
	char := '0'
	buf := bytes.NewBufferString("")

	for pos, char = range bits {
		if skipping_zeros {
			seq_start = pos
			if char == '0' {
				continue
			} else {
				skipping_zeros = false
			}
		}
		if prev != char {
			seq_len = pos - seq_start
			min_ = min(min_, seq_len)
			max_ = max(max_, seq_len)
			seq_start = pos
		}
		prev = char
	}
	if prev == '1' {
		seq_len = pos - seq_start + 1
		min_ = min(min_, seq_len)
		max_ = max(max_, seq_len)
	}
	skipping_zeros = true
	prev = '1'
	seq_start = 0
	seq_len = 0

	if min_ == math.MaxInt64 {
		panic("No ones found: " + bits)
	}
	if min_*7 == max_ {
		max_ = 3 * min_
	}
	if min_ == max_ {
		max_ = 3 * min_
	}
	for pos, char := range bits {
		if skipping_zeros {
			seq_start = pos
			if char == '0' {
				continue
			} else {
				skipping_zeros = false
			}
		}
		if prev != char {
			if prev == '1' {
				seq_len = pos - seq_start
				if seq_len == min_ {
					buf.WriteString(".")
				} else if seq_len == max_ {
					buf.WriteString("-")
				} else {
					fmt.Println("Ones length: ", seq_len)
					panic("Unrecognised ones length")
				}
			}
			if prev == '0' {
				seq_len = pos - seq_start
				if seq_len == min_ {
					// Nothing, just next part of string
				} else if seq_len == max_ {
					buf.WriteString(" ")
				} else if seq_len == min_*7 {
					buf.WriteString("   ")
				} else {
					fmt.Println("seq: ", bits)
					fmt.Println("Zeros length: ", seq_len, min_, max_)
					panic("Unrecognised zeros length")
				}
			}
			seq_start = pos
		}
		prev = char
	}

	seq_len = pos - seq_start + 1
	if prev == '1' {
		if seq_len == min_ {
			buf.WriteString(".")
		} else if seq_len == max_ {
			buf.WriteString("-")
		} else {
			fmt.Println("Ones length: ", seq_len)
			panic("Unrecognised ones length")
		}
	}
	if prev == '0' {
		if seq_len == min_ {
			// Nothing, just next part of string
		} else if seq_len == max_ {
			buf.WriteString(" ")
		} else if seq_len == min_*7 {
			buf.WriteString("   ")
		}
	}
	return buf.String()
}

func DecodeMorse(morseCode string) string {
	char := '0'
	morse_seq := ""
	res := ""
	space_count := 0
	var morse_map = map[string]string{
		".-":     "A",
		"-...":   "B",
		"-.-.":   "C",
		"-..":    "D",
		".":      "E",
		"..-.":   "F",
		"--.":    "G",
		"....":   "H",
		"..":     "I",
		".---":   "J",
		"-.-":    "K",
		".-..":   "L",
		"--":     "M",
		"-.":     "N",
		"---":    "O",
		".--.":   "P",
		"--.-":   "Q",
		".-.":    "R",
		"...":    "S",
		"-":      "T",
		"..-":    "U",
		"...-":   "V",
		".--":    "W",
		"-..-":   "X",
		"-.--":   "Y",
		"--..":   "Z",
		".-.-.-": ".",
	}
	is_char := true
	for _, char = range morseCode {
		fmt.Println("seq: " + morse_seq)
		fmt.Println("char: " + string(char))
		if char != ' ' {
			if !is_char {
				if space_count == 3 {
					res = res + " "
				}
				space_count = 0
				is_char = true
			}
			morse_seq = morse_seq + string(char)
		} else {
			if is_char {
				decoded_char, ok := morse_map[morse_seq]
				if !ok {
					panic("Unknown sequence: " + morse_seq)
				}
				res = res + decoded_char
				is_char = false
			}
			morse_seq = ""
			space_count = space_count + 1
		}
	}
	if len(morse_seq) > 0 {
		decoded_char, ok := morse_map[morse_seq]
		if !ok {
			panic("Unknown sequence: " + morse_seq)
		}
		res = res + decoded_char
	}
	return res
}

func main() {
	// DecodeMorse();
	res := DecodeBits("1100110011001100000011000000111111001100111111001111110000000000000011001111110011111100111111000000110011001111110000001111110011001100000011")
	fmt.Println(res)
	demorsed := DecodeMorse(res)
	fmt.Println(demorsed)

	res = DecodeBits("00011100010101010001000000011101110101110001010111000101000111010111010001110101110000000111010101000101110100011101110111000101110111000111010000000101011101000111011101110001110101011100000001011101110111000101011100011101110001011101110100010101000000011101110111000101010111000100010111010000000111000101010100010000000101110101000101110001110111010100011101011101110000000111010100011101110111000111011101000101110101110101110")
	fmt.Println(res)
	demorsed = DecodeMorse(res)
	fmt.Println(demorsed)
}
