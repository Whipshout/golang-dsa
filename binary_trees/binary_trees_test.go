package binary_trees

import (
	"golang.org/x/exp/constraints"
	"reflect"
	"testing"
)

func TestBinaryTree_Delete(t *testing.T) {
	type args[T constraints.Ordered] struct {
		value T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		bt   BinaryTree[T]
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "Delete leaf node",
			bt:   BinaryTree[int]{Root: &Node[int]{Value: 5, Left: &Node[int]{Value: 3}, Right: &Node[int]{Value: 7}}},
			args: args[int]{value: 3},
			want: []int{5, 7},
		},
		{
			name: "Delete node with one child",
			bt:   BinaryTree[int]{Root: &Node[int]{Value: 5, Left: &Node[int]{Value: 3, Left: &Node[int]{Value: 1}}, Right: &Node[int]{Value: 7}}},
			args: args[int]{value: 3},
			want: []int{1, 5, 7},
		},
		{
			name: "Delete node with two children",
			bt:   BinaryTree[int]{Root: &Node[int]{Value: 5, Left: &Node[int]{Value: 3}, Right: &Node[int]{Value: 7, Left: &Node[int]{Value: 6}, Right: &Node[int]{Value: 8}}}},
			args: args[int]{value: 7},
			want: []int{3, 5, 6, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bt.Delete(tt.args.value)
			if got := tt.bt.InorderTraversal(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("After Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTree_InorderTraversal(t *testing.T) {
	type testCase[T constraints.Ordered] struct {
		name string
		bt   BinaryTree[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "Empty tree",
			bt:   BinaryTree[int]{},
			want: []int{},
		},
		{
			name: "Tree with one node",
			bt:   BinaryTree[int]{Root: &Node[int]{Value: 5}},
			want: []int{5},
		},
		{
			name: "Tree with multiple nodes",
			bt:   BinaryTree[int]{Root: &Node[int]{Value: 5, Left: &Node[int]{Value: 3, Left: &Node[int]{Value: 1}, Right: &Node[int]{Value: 4}}, Right: &Node[int]{Value: 7, Left: &Node[int]{Value: 6}, Right: &Node[int]{Value: 8}}}},
			want: []int{1, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bt.InorderTraversal(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTree_Insert(t *testing.T) {
	type args[T constraints.Ordered] struct {
		value T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		bt   BinaryTree[T]
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "Add into empty tree",
			bt:   BinaryTree[int]{},
			args: args[int]{value: 5},
			want: []int{5},
		},
		{
			name: "Add smaller value",
			bt:   BinaryTree[int]{Root: &Node[int]{Value: 5}},
			args: args[int]{value: 3},
			want: []int{3, 5},
		},
		{
			name: "Add larger value",
			bt:   BinaryTree[int]{Root: &Node[int]{Value: 5}},
			args: args[int]{value: 7},
			want: []int{5, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bt.Insert(tt.args.value)
			if got := tt.bt.InorderTraversal(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("After Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTree_Search(t *testing.T) {
	type args[T constraints.Ordered] struct {
		value T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		bt   BinaryTree[T]
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "Search in empty tree",
			bt:   BinaryTree[int]{},
			args: args[int]{value: 5},
			want: false,
		},
		{
			name: "Search existing value",
			bt:   BinaryTree[int]{Root: &Node[int]{Value: 5, Left: &Node[int]{Value: 3}, Right: &Node[int]{Value: 7}}},
			args: args[int]{value: 3},
			want: true,
		},
		{
			name: "Search non-existing value",
			bt:   BinaryTree[int]{Root: &Node[int]{Value: 5, Left: &Node[int]{Value: 3}, Right: &Node[int]{Value: 7}}},
			args: args[int]{value: 6},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bt.Search(tt.args.value); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type testCase[T constraints.Ordered] struct {
		name string
		want *BinaryTree[T]
	}
	tests := []testCase[int]{
		{
			name: "Create new binary tree",
			want: &BinaryTree[int]{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New[int](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteNode(t *testing.T) {
	type args[T constraints.Ordered] struct {
		node  *Node[T]
		value T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		args args[T]
		want *Node[T]
	}
	tests := []testCase[int]{
		{
			name: "Delete leaf node",
			args: args[int]{
				node:  &Node[int]{Value: 5, Left: &Node[int]{Value: 3}, Right: &Node[int]{Value: 7}},
				value: 3,
			},
			want: &Node[int]{Value: 5, Right: &Node[int]{Value: 7}},
		},
		{
			name: "Delete node with one child",
			args: args[int]{
				node:  &Node[int]{Value: 5, Left: &Node[int]{Value: 3, Left: &Node[int]{Value: 1}}, Right: &Node[int]{Value: 7}},
				value: 3,
			},
			want: &Node[int]{Value: 5, Left: &Node[int]{Value: 1}, Right: &Node[int]{Value: 7}},
		},
		{
			name: "Delete node with two children",
			args: args[int]{
				node:  &Node[int]{Value: 5, Left: &Node[int]{Value: 3}, Right: &Node[int]{Value: 7, Left: &Node[int]{Value: 6}, Right: &Node[int]{Value: 8}}},
				value: 7,
			},
			want: &Node[int]{Value: 5, Left: &Node[int]{Value: 3}, Right: &Node[int]{Value: 8, Left: &Node[int]{Value: 6}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteNode(tt.args.node, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMin(t *testing.T) {
	type args[T constraints.Ordered] struct {
		node *Node[T]
	}
	type testCase[T constraints.Ordered] struct {
		name string
		args args[T]
		want *Node[T]
	}
	tests := []testCase[int]{
		{
			name: "Find min in left-skewed tree",
			args: args[int]{node: &Node[int]{Value: 5, Left: &Node[int]{Value: 3, Left: &Node[int]{Value: 1}}}},
			want: &Node[int]{Value: 1},
		},
		{
			name: "Find min in balanced tree",
			args: args[int]{node: &Node[int]{Value: 5, Left: &Node[int]{Value: 3}, Right: &Node[int]{Value: 7}}},
			want: &Node[int]{Value: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inorderTraversal(t *testing.T) {
	type args[T constraints.Ordered] struct {
		node   *Node[T]
		result *[]T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "Inorder traversal of balanced tree",
			args: args[int]{
				node:   &Node[int]{Value: 5, Left: &Node[int]{Value: 3}, Right: &Node[int]{Value: 7}},
				result: &[]int{},
			},
			want: []int{3, 5, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inorderTraversal(tt.args.node, tt.args.result)
			if !reflect.DeepEqual(*tt.args.result, tt.want) {
				t.Errorf("inorderTraversal() result = %v, want %v", *tt.args.result, tt.want)
			}
		})
	}
}

func Test_insert(t *testing.T) {
	type args[T constraints.Ordered] struct {
		node  *Node[T]
		value T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		args args[T]
		want *Node[T]
	}
	tests := []testCase[int]{
		{
			name: "Add into empty tree",
			args: args[int]{node: nil, value: 5},
			want: &Node[int]{Value: 5},
		},
		{
			name: "Insert smaller value",
			args: args[int]{node: &Node[int]{Value: 5}, value: 3},
			want: &Node[int]{Value: 5, Left: &Node[int]{Value: 3}},
		},
		{
			name: "Insert larger value",
			args: args[int]{node: &Node[int]{Value: 5}, value: 7},
			want: &Node[int]{Value: 5, Right: &Node[int]{Value: 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.node, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_search(t *testing.T) {
	type args[T constraints.Ordered] struct {
		node  *Node[T]
		value T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "Search in empty tree",
			args: args[int]{node: nil, value: 5},
			want: false,
		},
		{
			name: "Search existing value",
			args: args[int]{node: &Node[int]{Value: 5, Left: &Node[int]{Value: 3}, Right: &Node[int]{Value: 7}}, value: 3},
			want: true,
		},
		{
			name: "Search non-existing value",
			args: args[int]{node: &Node[int]{Value: 5, Left: &Node[int]{Value: 3}, Right: &Node[int]{Value: 7}}, value: 6},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.node, tt.args.value); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
