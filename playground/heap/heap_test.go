package playground

import (
	"reflect"
	"testing"
)

func TestMinHeapSize(t *testing.T) {
	testCases := []struct {
		name     string
		heap     *MinHeap
		expected int
	}{
		{name: "1", heap: &MinHeap{array: []int{1}}, expected: 1},
		{name: "2", heap: &MinHeap{}, expected: 0},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.heap.Size()
			want := tt.expected

			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		})
	}
}

func TestMinHeapPeek(t *testing.T) {
	testCases := []struct {
		name     string
		heap     *MinHeap
		expected int
	}{
		{name: "1", heap: &MinHeap{array: []int{77}}, expected: 77},
		{name: "2", heap: &MinHeap{}, expected: 0},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.heap.Peek()
			want := tt.expected

			if got != want {
				t.Errorf("got %d want %d; error %v", got, want, err)
			}
		})
	}
}

func TestMinHeapInsert(t *testing.T) {
	testCases := []struct {
		name     string
		heap     *MinHeap
		key      int
		expected int
	}{
		{name: "1", heap: &MinHeap{array: []int{2, 3}}, key: 1, expected: 1},
		{name: "2", heap: &MinHeap{array: []int{66, 77}}, key: 88, expected: 66},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.heap.Insert(tt.key)

			got, err := tt.heap.Peek()
			want := tt.expected

			if got != want {
				t.Errorf("got %d want %d; error %v", got, want, err)
			}
		})
	}
}

func TestMinHeapExtractMin(t *testing.T) {
	testCases := []struct {
		name     string
		heap     *MinHeap
		expected int
	}{
		{name: "1", heap: &MinHeap{array: []int{8, 14, 12, 9, 16}}, expected: 8},
		{name: "2", heap: &MinHeap{array: []int{}}, expected: 0},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.heap.ExtractMin()
			want := tt.expected

			if got != want {
				t.Errorf("got %d want %d; error %v", got, want, err)
			}
		})
	}
}

func TestMinHeapHeapifyDown(t *testing.T) {
	testCases := []struct {
		name     string
		heap     *MinHeap
		expected []int
	}{
		{
			name:     "1",
			heap:     &MinHeap{array: []int{10, 15, 20, 17, 25}},
			expected: []int{15, 17, 20, 25},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.heap.array[0] = tt.heap.array[len(tt.heap.array)-1]
			tt.heap.array = tt.heap.array[:len(tt.heap.array)-1]

			tt.heap.HeapifyDown(0)

			if !reflect.DeepEqual(tt.heap.array, tt.expected) {
				t.Errorf("got %v want %v", tt.heap.array, tt.expected)
			}
		})
	}
}

func TestMinHeapHeapifyUp(t *testing.T) {
	testCases := []struct {
		name     string
		heap     *MinHeap
		expected []int
	}{
		{
			name:     "1",
			heap:     &MinHeap{[]int{10, 15, 20, 17, 8}},
			expected: []int{8, 10, 20, 17, 15},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.heap.HeapifyUp(len(tt.heap.array) - 1)

			if !reflect.DeepEqual(tt.heap.array, tt.expected) {
				t.Errorf("got %v want %v", tt.heap.array, tt.expected)
			}
		})
	}
}
