package trees

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args[K comparable, V any] struct {
		rootKey   K
		rootValue V
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
		want *Tree[K, V]
	}
	tests := []testCase[string, string]{
		{
			name: "Create new tree",
			args: args[string, string]{rootKey: "A", rootValue: "root_value"},
			want: &Tree[string, string]{Root: &Node[string, string]{key: "A", value: "root_value"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.rootKey, tt.args.rootValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_AddChild(t *testing.T) {
	type args[K comparable, V any] struct {
		childKey   K
		childValue V
	}
	type testCase[K comparable, V any] struct {
		name string
		n    *Node[K, V]
		args args[K, V]
		want *Node[K, V]
	}
	tests := []testCase[string, string]{
		{
			name: "Add child to node",
			n:    &Node[string, string]{key: "A", value: "root_value"},
			args: args[string, string]{childKey: "B", childValue: "valueB"},
			want: &Node[string, string]{key: "B", value: "valueB"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.AddChild(tt.args.childKey, tt.args.childValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddChild() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_Display(t *testing.T) {
	type args struct {
		level int
	}
	type testCase[K comparable, V any] struct {
		name string
		n    *Node[K, V]
		args args
	}
	tests := []testCase[string, string]{
		{
			name: "Display tree",
			n: func() *Node[string, string] {
				root := &Node[string, string]{key: "A", value: "root_value"}
				root.AddChild("B", "valueB")
				return root
			}(),
			args: args{level: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.Display(tt.args.level)
		})
	}
}

func TestNode_Find(t *testing.T) {
	type args[K comparable] struct {
		key K
	}
	type testCase[K comparable, V any] struct {
		name string
		n    *Node[K, V]
		args args[K]
		want *Node[K, V]
	}
	tests := []testCase[string, string]{
		{
			name: "Find existing node",
			n: func() *Node[string, string] {
				root := &Node[string, string]{key: "A", value: "root_value"}
				root.AddChild("B", "valueB")
				return root
			}(),
			args: args[string]{key: "B"},
			want: &Node[string, string]{key: "B", value: "valueB"},
		},
		{
			name: "Find non-existing node",
			n:    &Node[string, string]{key: "A", value: "root_value"},
			args: args[string]{key: "C"},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Find(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_RemoveChild(t *testing.T) {
	type args[K comparable] struct {
		key K
	}
	type testCase[K comparable, V any] struct {
		name string
		n    *Node[K, V]
		args args[K]
		want bool
	}
	tests := []testCase[string, string]{
		{
			name: "Remove existing child",
			n: func() *Node[string, string] {
				root := &Node[string, string]{key: "A", value: "root_value"}
				root.AddChild("B", "valueB")
				return root
			}(),
			args: args[string]{key: "B"},
			want: true,
		},
		{
			name: "Remove non-existing child",
			n:    &Node[string, string]{key: "A", value: "root_value"},
			args: args[string]{key: "C"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.RemoveChild(tt.args.key); got != tt.want {
				t.Errorf("RemoveChild() = %v, want %v", got, tt.want)
			}
		})
	}
}
