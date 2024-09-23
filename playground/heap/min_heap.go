/*
Min Heap:
A Heap uses a binary tree structure to hold its nodes where the smallest node is
the root node (or largest for a max heap). Instead of implementing the heap as
a node with a left and right pointer, we use an array to store their values.

	This array:
	[3, 4, 8, 9, 7, 10, 9]

	Is the same as this tree:
             3
		   /   \
		  4     8
		 / \   / \
		9   7 10  9

To access family members of the tree we can use these formulas:
	Find the parent to a node:
	parent_index = (current_index - 1) / 2

	Find the left child of a node:
	left_index = 2 * current_index + 1

	Find the right child of a node:
	right_index = 2 * current_index + 2
*/

package playground

import "errors"

type MinHeap struct {
	array []int
}

/*
Insert - add a new value to the heap while maintaining the heap properties.

Inserting a node into a heap always goes into the next available spot and if it
breaks the rules of the heap, we bubble it to the top.
*/
func (h *MinHeap) Insert(key int) {
	h.array = append(h.array, key)

	h.HeapifyUp(len(h.array) - 1)
}

/*
ExtractMin - remove the root of the heap while maintaining the heap properties.

When we need to remove the root node, we first have to swap the roots
value with the last node added. The heap might not be in the right order
so we need to bubble the node down until the heap order is restored.
*/
func (h *MinHeap) ExtractMin() (int, error) {
	if len(h.array) == 0 {
		return 0, errors.New("extracting from an empty heap")
	}

	root := h.array[0]

	// Replace root with the last element before calling HeapifyDown
	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]

	h.HeapifyDown(0)

	return root, nil
}

// Peek - return the root element without removing it.
func (h *MinHeap) Peek() (int, error) {
	if len(h.array) == 0 {
		return 0, errors.New("peek on an empty heap")
	}

	return h.array[0], nil
}

// Size - return the size of all the elements in the heap.
func (h *MinHeap) Size() int {
	return len(h.array)
}

// HeapifyUp - maintain the heap property after insertion (bubble up).
func (h *MinHeap) HeapifyUp(i int) {
	for i > 0 {
		// Find the parent of the current node
		parent := (i - 1) / 2

		// Swap nodes if parent violates the heap property
		if h.array[parent] > h.array[i] {
			h.array[parent], h.array[i] = h.array[i], h.array[parent]
			i = parent
		} else {
			break
		}
	}
}

// HeapifyDown - maintain the heap property after root extraction (bubble down).
func (h *MinHeap) HeapifyDown(i int) {
	last := len(h.array) - 1

	for {
		// Find the left and right nodes
		left, right := 2*i+1, 2*i+2
		smallest := i

		// Check if the left child exists and is smaller than the current node
		if left <= last && h.array[left] < h.array[smallest] {
			smallest = left
		}

		// Check if the right child exists and is smaller than the current node
		if right <= last && h.array[right] < h.array[smallest] {
			smallest = right
		}

		// Check if the current node is the smallest
		if smallest == i {
			break
		}

		// Swap the current with the smallest and continue down the tree
		h.array[smallest], h.array[i] = h.array[i], h.array[smallest]
		i = smallest
	}
}
