// https://leetcode.com/problems/relative-ranks/description/

package heap

import (
	"container/heap"
	"fmt"
)

// Max Heap Implementation
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func FindRelativeRanks(score []int) []string {
	answer := make([]string, len(score))
	h := &IntHeap{}
	heap.Init(h)

	for _, s := range score {
		heap.Push(h, s)
	}

	place := 1

	// Pop each element off the priority queue
	for h.Len() > 0 {
		rank := heap.Pop(h)

		// Find that element in the scores slice
		for i, s := range score {
			if rank == s {
				if place == 1 {
					answer[i] = "Gold Medal"
				} else if place == 2 {
					answer[i] = "Silver Medal"
				} else if place == 3 {
					answer[i] = "Bronze Medal"
				} else {
					answer[i] = fmt.Sprintf("%d", place)
				}

				break
			}
		}

		place++
	}

	return answer
}
