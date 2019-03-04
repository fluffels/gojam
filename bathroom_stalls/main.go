package main

import (
	"fmt"
)

func solve(N, K int) (int, int) {
	counts := map[int]int{
		N: 1,
	}
	for {
		fmt.Println(counts)
		size, count := 0, 0
		for s, c := range counts {
			if (c > 0) && (s > size) {
				size = s
				count = c
			}
		}
		K -= count
		counts[size] = 0
		newSize := size - 1
		left := newSize / 2
		var right int
		if newSize > 1 {
			if newSize%2 == 0 {
				right = left
			} else {
				right = left + 1
			}
		} else {
			right = 0
		}
		if K <= 0 {
			return left, right
		} else {
			if left > 0 {
				c := counts[left]
				counts[left] = c + count
			}
			if right > 0 {
				c := counts[right]
				counts[right] = c + count
			}
		}
	}
}

func main() {
	var T int
	fmt.Scan(&T)
	for t := 1; t <= T; t++ {
		var N, K int
		fmt.Scan(&N, &K)
		left, right := solve(N, K)
		fmt.Printf("Case #%d: %d %d\n", t, left, right)
	}
}
