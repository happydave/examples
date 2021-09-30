package main

import (
	"fmt"
	"os"
)

// x = fib(n-a) + fib(n-2)

// store generated fib sequence numbers to avoid duplicating work
var seq = []int{0,1} 

func main() {
												// expected result
	fmt.Printf("%d: %d \n", 0, fib(0))			// 0
	fmt.Printf("%d: %d \n", -1, fib(-1))		// -1
	fmt.Printf("%d: %d \n", 12, fib(12))		// 144
	fmt.Printf("%d: %d \n", 1, fib(1))			// 1
	fmt.Printf("%d: %d \n", 7, fib(7))			// 13
	fmt.Printf("%d: %d \n", -13, fib(-13))		// -1
	fmt.Printf("%d: %d \n", 92, fib(92))		// 7540113804746346429
	fmt.Printf("%d: %d \n", 93, fib(93))		// -6246583658587674878 <== rolled the int
	fmt.Printf("%d: %d \n", 95, fib(95))
	//fmt.Printf("%d: %d \n", 65537, fib(65537))  // 4169236831028050205 is incorrect
}

// fib(index) returns fib(index-1) + fib(index-2)
func fib(n int) int {

	if n == 0 { return seq[0] }
	if n == 1 { return seq[1] }
	if n < 0  { return -1 }
	
	// Generate numbers to requested index if needed
	if len(seq) < n {
		for i := len(seq); i <= n; i++{
			seq = append(seq, seq[i-1] + seq[i-2])
			
			if seq[i] < 0 {
				// ERROR: rolled the int
				fmt.Printf("ERROR: Int roll at %d (%d + %d returned %d) \n", i, seq[i-2], seq[i-1], seq[i])
				os.Exit(1)
			}
		}
	}
	
	return seq[n-1] + seq[n-2]
}
