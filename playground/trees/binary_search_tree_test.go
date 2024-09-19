package playground

import "testing"

func TestBSTInsert(t *testing.T) {
	tree := &BinaryTree{}

	tree.BSTInsert(2)
	tree.BSTInsert(1)
	tree.BSTInsert(3)
	tree.BSTInsert(4)
	tree.BSTInsert(6)
	tree.BSTInsert(8)
	tree.BSTInsert(7)
	tree.BSTInsert(9)

	TreeDisplay(tree.root, 0)
}

func TestBSTSearch(t *testing.T) {
	tree := &BinaryTree{}

	tree.BSTInsert(2)
	tree.BSTInsert(1)
	tree.BSTInsert(3)

	testCases := []struct {
		name     string
		target   int
		expected bool
	}{
		{name: "1", target: 3, expected: true},
		{name: "2", target: 17, expected: false},
		{name: "3", target: 2, expected: true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := tree.BSTSearch(tt.target)
			want := tt.expected

			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}

func TestBSTTraversalInOrder(t *testing.T) {
	tree := &BinaryTree{}

	tree.BSTInsert(5)
	tree.BSTInsert(2)
	tree.BSTInsert(6)
	tree.BSTInsert(1)
	tree.BSTInsert(3)

	tree.BSTTraversalInOrder()
}

func TestBSTTraversalPreOrder(t *testing.T) {
	tree := &BinaryTree{}

	tree.BSTInsert(5)
	tree.BSTInsert(2)
	tree.BSTInsert(6)
	tree.BSTInsert(1)
	tree.BSTInsert(3)

	tree.BSTTraveralPreOrder()
}

func TestBSTTraversalPostOrder(t *testing.T) {
	tree := &BinaryTree{}

	tree.BSTInsert(5)
	tree.BSTInsert(2)
	tree.BSTInsert(6)
	tree.BSTInsert(1)
	tree.BSTInsert(3)

	tree.BSTTraversalPostOrder()
}

func TestBSTTraversalLevelOrder(t *testing.T) {
	tree := &BinaryTree{}

	tree.BSTInsert(5)
	tree.BSTInsert(2)
	tree.BSTInsert(6)
	tree.BSTInsert(1)
	tree.BSTInsert(3)

	tree.BSTTraversalLevelOrder()
}

func TestBSTRemove(t *testing.T) {
	tree := &BinaryTree{}

	tree.BSTInsert(5)
	tree.BSTInsert(2)
	tree.BSTInsert(6)
	tree.BSTInsert(1)
	tree.BSTInsert(3)

	tree.Remove(5)

	TreeDisplay(tree.root, 0)
}
