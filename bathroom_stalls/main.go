package main

import (
	"container/heap"
	"fmt"
)

type Partition struct {
	size  int
	count int
}

type PartitionHeap []Partition

func (h PartitionHeap) Len() int           { return len(h) }
func (h PartitionHeap) Less(i, j int) bool { return h[i].size < h[j].size }
func (h PartitionHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PartitionHeap) Push(x interface{}) {
	*h = append(*h, x.(Partition))
}

func (h *PartitionHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func solve(N, K int) (int, int) {
	h := &PartitionHeap{{N, 1}}
	heap.Init(h)
	for {
		p := h.Pop().(Partition)
		K -= p.count
		newSize := p.size - 1
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
			leftFound, rightFound := false, false
			for _, pPrime := range *h {
				if (left > 0) && (pPrime.size == left) {
					pPrime.count += p.count
					leftFound = true
				}
				if (right > 0) && (pPrime.size == right) {
					pPrime.count += p.count
					rightFound = true
				}
			}
			if (left > 0) && (leftFound == false) {
				newP := Partition{left, p.count}
				h.Push(newP)
			}
			if (right > 0) && (rightFound == false) {
				newP := Partition{right, p.count}
				h.Push(newP)
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
