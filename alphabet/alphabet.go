// imagine a dial with A-Z on it (A adjacent to Z)
// it takes one tick to move the dial one character

// challenge: write a function that returns the total number of ticks per character string
//  count ticks between each character in order
//  always use the shortest path
//  dial starts on A
// input: string of characters (length <= 20 characters, only alpha characters)
// output: total number of ticks

// Example: BYCD
//    A->B = 1 ticks
//    B->Y = 3 ticks
//    Y->C = 4 ticks
//    Total: 8 ticks

package alphabet

import (
	"strings"
)

// go test -v *.go

func countTicks(input string) int {
	// int('A') == 65 // fmt.Printf("A:%d\n", int('A'))
	// int('Z') == 90 // fmt.Printf("Z:%d\n", int('Z'))
	current := int('A')
	count := 0
	input = strings.ToUpper(input)

	for _, c := range input {
		ci := int(c)
		diff := ci - current

		if diff < 0 {
			diff = diff + 26
		}

		if diff > 13 {
			diff = 26 - diff
		}

		count = count + diff
		current = ci
		//fmt.Printf("%c (%d) %d\n", c, ci, diff)
	}

	return count
}
