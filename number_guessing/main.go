package main

import (
	"fmt"
)

func solve(A, B, N int) {
	var guess int
	var response string
	for {
		// NOTE(jan): codejam uses GO 1.7 which has no math.Round
		add := (B - A) / 2
		if add == 0 {
			add = 1
		}
		guess = A + add
		fmt.Println(guess)
		fmt.Scan(&response)
		switch response {
		case "CORRECT":
			return
		case "TOO_SMALL":
			A = guess
		case "TOO_BIG":
			B = guess
		case "WRONG_ANSWER":
			return
		}
	}
}

func main() {
	var T int
	fmt.Scan(&T)
	for t := 0; t < T; t++ {
		var A, B, N int
		fmt.Scan(&A, &B, &N)
		solve(A, B, N)
	}
}
