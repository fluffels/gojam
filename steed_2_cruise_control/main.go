package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)
	for t := 1; t <= T; t++ {
		var D, N int
		fmt.Scan(&D, &N)
		var tMax float64
		tMax = 0.0
		for n := 0; n < N; n++ {
			var k, s int
			fmt.Scan(&k, &s)
			t := float64(D-k) / float64(s)
			if t >= tMax {
				tMax = t
			}
		}
		s := float64(D) / tMax
		fmt.Printf("Case #%d: %f\n", t, s)
	}
}
