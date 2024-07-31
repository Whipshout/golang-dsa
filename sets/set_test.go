package sets

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		want *Set[T]
	}
	tests := []testCase[int]{
		{
			name: "Create new int Set",
			want: &Set[int]{},
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

func TestSet_Add(t *testing.T) {
	type args[T comparable] struct {
		element T
	}
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		args args[T]
		want Set[T]
	}
	tests := []testCase[int]{
		{
			name: "Add to empty set",
			s:    make(Set[int]),
			args: args[int]{element: 1},
			want: Set[int]{1: struct{}{}},
		},
		{
			name: "Add to non-empty set",
			s:    Set[int]{2: struct{}{}, 3: struct{}{}},
			args: args[int]{element: 1},
			want: Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
		},
		{
			name: "Add existing element",
			s:    Set[int]{1: struct{}{}, 2: struct{}{}},
			args: args[int]{element: 1},
			want: Set[int]{1: struct{}{}, 2: struct{}{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Add(tt.args.element)
			if !reflect.DeepEqual(tt.s, tt.want) {
				t.Errorf("After Add(%v), got %v, want %v", tt.args.element, tt.s, tt.want)
			}
		})
	}
}

func TestSet_Clear(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		s    Set[T]
	}
	tests := []testCase[int]{
		{
			name: "Clear empty set",
			s:    make(Set[int]),
		},
		{
			name: "Clear set with one element",
			s:    Set[int]{1: struct{}{}},
		},
		{
			name: "Clear set with multiple elements",
			s:    Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Clear()
			if len(tt.s) != 0 {
				t.Errorf("After Clear(), set is not empty. Got %v elements", len(tt.s))
			}
		})
	}
}

func TestSet_Contains(t *testing.T) {
	type args[T comparable] struct {
		element T
	}
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "Empty set doesn't contain element",
			s:    make(Set[int]),
			args: args[int]{element: 1},
			want: false,
		},
		{
			name: "Set contains the element",
			s:    Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			args: args[int]{element: 2},
			want: true,
		},
		{
			name: "Set doesn't contain the element",
			s:    Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			args: args[int]{element: 4},
			want: false,
		},
		{
			name: "Set with one element contains that element",
			s:    Set[int]{5: struct{}{}},
			args: args[int]{element: 5},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Contains(tt.args.element); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Difference(t *testing.T) {
	type args[T comparable] struct {
		other Set[T]
	}
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		args args[T]
		want *Set[T]
	}
	tests := []testCase[int]{
		{
			name: "Set with different values",
			s:    Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			args: args[int]{other: Set[int]{1: struct{}{}}},
			want: &Set[int]{2: struct{}{}, 3: struct{}{}},
		},
		{
			name: "Set without different values",
			s:    Set[int]{1: struct{}{}, 2: struct{}{}},
			args: args[int]{other: Set[int]{1: struct{}{}, 2: struct{}{}}},
			want: &Set[int]{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Difference(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Difference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Intersection(t *testing.T) {
	type args[T comparable] struct {
		other Set[T]
	}
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		args args[T]
		want *Set[T]
	}
	tests := []testCase[int]{
		{
			name: "Set with intersectional values",
			s:    Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			args: args[int]{other: Set[int]{1: struct{}{}}},
			want: &Set[int]{1: struct{}{}},
		},
		{
			name: "Set without intersectional values",
			s:    Set[int]{1: struct{}{}, 2: struct{}{}},
			args: args[int]{other: Set[int]{3: struct{}{}, 4: struct{}{}}},
			want: &Set[int]{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Intersection(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Union(t *testing.T) {
	type args[T comparable] struct {
		other Set[T]
	}
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		args args[T]
		want *Set[T]
	}
	tests := []testCase[int]{
		{
			name: "Set with union values",
			s:    Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			args: args[int]{other: Set[int]{4: struct{}{}}},
			want: &Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Union(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Remove(t *testing.T) {
	type args[T comparable] struct {
		element T
	}
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		args args[T]
		want Set[T]
	}
	tests := []testCase[int]{
		{
			name: "Remove from empty set",
			s:    make(Set[int]),
			args: args[int]{element: 1},
			want: make(Set[int]),
		},
		{
			name: "Remove existing element",
			s:    Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			args: args[int]{element: 2},
			want: Set[int]{1: struct{}{}, 3: struct{}{}},
		},
		{
			name: "Remove non-existing element",
			s:    Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			args: args[int]{element: 4},
			want: Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
		},
		{
			name: "Remove from set with one element",
			s:    Set[int]{5: struct{}{}},
			args: args[int]{element: 5},
			want: make(Set[int]),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Remove(tt.args.element)
			if !reflect.DeepEqual(tt.s, tt.want) {
				t.Errorf("After Remove(%v), got %v, want %v", tt.args.element, tt.s, tt.want)
			}
		})
	}
}

func TestSet_Size(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		want int
	}
	tests := []testCase[int]{
		{
			name: "Size empty set",
			s:    make(Set[int]),
			want: 0,
		},
		{
			name: "Size set with one element",
			s:    Set[int]{1: struct{}{}},
			want: 1,
		},
		{
			name: "Size set with multiple elements",
			s:    Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_ToSlice(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "Empty set",
			s:    make(Set[int]),
			want: []int{},
		},
		{
			name: "Set with one element",
			s:    Set[int]{5: struct{}{}},
			want: []int{5},
		},
		{
			name: "Set with multiple elements",
			s:    Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("ToSlice() returned slice with length %v, want %v", len(got), len(tt.want))
			}
			uniqueCheck := make(map[int]bool)
			for _, v := range got {
				if uniqueCheck[v] {
					t.Errorf("ToSlice() returned slice with duplicate element: %v", v)
				}
				uniqueCheck[v] = true
			}
		})
	}
}
