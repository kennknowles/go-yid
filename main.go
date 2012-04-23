
package main

import "fmt"
import "yid"
import "time"

// Big enough for N from 1 to 250
var input = make([]string, 1, 499)

func main() {
	// Initialize the whole array
	for idx, _ := range input {
		if (idx % 2) == 0 {
			input[idx] = "+"
		} else {
			input[idx] = "one"
		}
	}

	// Take slices
	for n := 1 ; n <= 250 ; n++ {
		input = input[0:2*n-1]
		start := time.Now()
		fmt.Printf("%d,", (len(input)+1)/2)
		yid.Russ_cox_exponential.Recognize(input)
		fmt.Printf("%.3f\n", time.Now().Sub(start).Seconds())
	}
}
