package playground

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	var root *TreeNode

	root = root.Insert(2)
	root.Insert(1)
	root.Insert(3)
	root.Insert(4)
	root.Insert(6)
	root.Insert(8)
	root.Insert(7)
	root.Insert(9)

	TreeDisplay(root, 0)
}

func TestInsertDuplicate(t *testing.T) {
	var root *TreeNode

	root = root.Insert(2)
	root.Insert(3)
	root.Insert(4)
	root.Insert(4)

	TreeDisplay(root, 0)
}

func TestSearch(t *testing.T) {
	var root *TreeNode

	root = root.Insert(2)
	root.Insert(1)
	root.Insert(3)

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
			got := root.Search(tt.target)
			want := tt.expected

			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}

func TestInOrderTraversal(t *testing.T) {
	testCases := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name: "1",
			root: &TreeNode{data: 5,
				left: &TreeNode{data: 2,
					left: &TreeNode{data: 1,
						left: nil, right: nil},
					right: &TreeNode{data: 3,
						left: nil, right: nil}},
				right: &TreeNode{data: 6,
					left: nil, right: nil}},
			expected: []int{1, 2, 3, 5, 6},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.root.InOrderTraversal()
			want := tt.expected

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}

func TestPreOrderTraversal(t *testing.T) {
	var root *TreeNode

	root = root.Insert(5)
	root.Insert(2)
	root.Insert(6)
	root.Insert(1)
	root.Insert(3)

	root.PreOrderTraversal()
}

func TestPostOrderTraversal(t *testing.T) {
	var root *TreeNode

	root = root.Insert(5)
	root.Insert(2)
	root.Insert(6)
	root.Insert(1)
	root.Insert(3)

	root.PostOrderTraversal()
}

func TestLevelOrderTraversal(t *testing.T) {
	var root *TreeNode

	root = root.Insert(5)
	root.Insert(2)
	root.Insert(6)
	root.Insert(1)
	root.Insert(3)

	root.LevelOrderTraversal()
}

func TestRemove(t *testing.T) {
	var root *TreeNode

	root = root.Insert(5)
	root.Insert(2)
	root.Insert(6)
	root.Insert(1)
	root.Insert(3)

	root.Remove(2)

	TreeDisplay(root, 0)
}

func TestBalance(t *testing.T) {
	// Skewed BST
	var root *TreeNode

	root = root.Insert(4)
	root.Insert(3)
	root.Insert(5)
	root.Insert(2)
	root.Insert(6)
	root.Insert(1)
	root.Insert(7)

	balanced := root.Balance()

	TreeDisplay(balanced, 0)
}
