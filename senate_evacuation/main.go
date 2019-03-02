package main

import (
	"bytes"
	"fmt"
)

func sum(A []int) int {
	sum := 0
	for _, v := range A {
		sum += v
	}
	return sum
}

func solve(N int, P []int) []string {
	sumP := sum(P)
	plan := make([]string, sumP, sumP)
	for stepIdx := 0; sumP > 0; stepIdx++ {
		threshold := 0
		if sumP%2 == 0 {
			threshold = (sumP - 2) / 2
		} else {
			threshold = (sumP - 1) / 2
		}
		pick := func() byte {
			choice := -1
			for i, p := range P {
				if (choice == -1) && (p > 0) {
					choice = i
				} else if p > threshold {
					choice = i
					break
				}
			}
			P[choice] -= 1
			sumP -= 1
			return 'A' + byte(choice)
		}
		var buffer bytes.Buffer
		buffer.WriteByte(pick())
		if (sumP > 0) && (sumP%2 != 0) {
			buffer.WriteByte(pick())
		}
		plan[stepIdx] = buffer.String()
	}
	return plan
}

func main() {
	var T int
	fmt.Scanln(&T)
	for t := 1; t <= T; t++ {
		var N int
		fmt.Scanln(&N)
		P := make([]int, N, N)
		for n := 0; n < N; n++ {
			fmt.Scan(&(P[n]))
		}
		steps := solve(N, P)
		fmt.Printf("Case #%d: ", t)
		for _, step := range steps {
			fmt.Print(step, " ")
		}
		fmt.Println()
	}
}
