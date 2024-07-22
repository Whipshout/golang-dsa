package stacks

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		initialCapacity []int
	}
	type testCase[T any] struct {
		name string
		args args
		want *Stack[T]
	}
	tests := []testCase[bool]{
		{
			name: "New stack with default capacity",
			args: args{initialCapacity: []int{}},
			want: &Stack[bool]{
				elements: make([]bool, 0, 10),
				capacity: 10,
			},
		},
		{
			name: "New stack with custom capacity",
			args: args{initialCapacity: []int{20}},
			want: &Stack[bool]{
				elements: make([]bool, 0, 20),
				capacity: 20,
			},
		},
		{
			name: "New stack with zero capacity (should use default)",
			args: args{initialCapacity: []int{0}},
			want: &Stack[bool]{
				elements: make([]bool, 0, 10),
				capacity: 10,
			},
		},
		{
			name: "New stack with negative capacity (should use default)",
			args: args{initialCapacity: []int{-5}},
			want: &Stack[bool]{
				elements: make([]bool, 0, 10),
				capacity: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New[bool](tt.args.initialCapacity...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Clear(t *testing.T) {
	type testCase[T any] struct {
		name string
		s    *Stack[T]
		want *Stack[T]
	}
	tests := []testCase[bool]{
		{
			name: "Clear stack with elements",
			s: &Stack[bool]{
				elements: []bool{true, false, true},
				capacity: 10,
			},
			want: &Stack[bool]{
				elements: []bool{},
				capacity: 10,
			},
		},
		{
			name: "Clear stack without elements",
			s: &Stack[bool]{
				elements: []bool{},
				capacity: 10,
			},
			want: &Stack[bool]{
				elements: []bool{},
				capacity: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Clear()
			if !reflect.DeepEqual(tt.s, tt.want) {
				t.Errorf("Clear() = %v, want %v", tt.s, tt.want)
			}
		})
	}
}

func TestStack_IsEmpty(t *testing.T) {
	type testCase[T any] struct {
		name string
		s    *Stack[T]
		want bool
	}
	tests := []testCase[bool]{
		{
			name: "Stack is not empty",
			s: &Stack[bool]{
				elements: []bool{true, false, true},
				capacity: 10,
			},
			want: false,
		},
		{
			name: "Stack is empty",
			s: &Stack[bool]{
				elements: []bool{},
				capacity: 10,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Peek(t *testing.T) {
	type testCase[T any] struct {
		name    string
		s       *Stack[T]
		want    T
		wantErr bool
	}
	tests := []testCase[bool]{
		{
			name: "Peek stack with elements",
			s: &Stack[bool]{
				elements: []bool{true, false},
				capacity: 10,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Peek empty stack",
			s: &Stack[bool]{
				elements: []bool{},
				capacity: 10,
			},
			want:    true,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Peek()
			if (err != nil) != tt.wantErr {
				t.Errorf("Peek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Peek() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	type testCase[T any] struct {
		name    string
		s       *Stack[T]
		want    T
		wantErr bool
	}
	tests := []testCase[bool]{
		{
			name: "Pop stack with elements",
			s: &Stack[bool]{
				elements: []bool{true, false},
				capacity: 10,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Pop empty stack",
			s: &Stack[bool]{
				elements: []bool{},
				capacity: 10,
			},
			want:    true,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {
	type args[T any] struct {
		element T
	}
	type testCase[T any] struct {
		name    string
		s       *Stack[T]
		args    args[T]
		want    *Stack[T]
		wantErr bool
	}
	tests := []testCase[bool]{
		{
			name: "Push element to empty stack",
			s: &Stack[bool]{
				elements: []bool{},
				capacity: 10,
			},
			args: args[bool]{element: true},
			want: &Stack[bool]{
				elements: []bool{true},
				capacity: 10,
			},
			wantErr: false,
		},
		{
			name: "Push element to stack with elements",
			s: &Stack[bool]{
				elements: []bool{true, false},
				capacity: 10,
			},
			args: args[bool]{element: true},
			want: &Stack[bool]{
				elements: []bool{true, false, true},
				capacity: 10,
			},
			wantErr: false,
		},
		{
			name: "Push element to stack with no capacity",
			s: &Stack[bool]{
				elements: []bool{true, false},
				capacity: 2,
			},
			args: args[bool]{element: true},
			want: &Stack[bool]{
				elements: []bool{true, false, true},
				capacity: 10,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.s.Push(tt.args.element)
			if (err != nil) != tt.wantErr {
				t.Errorf("Push() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(tt.s, tt.want) {
				t.Errorf("Push() got = %v, want %v", tt.s, tt.want)
			}
		})
	}
}

func TestStack_Size(t *testing.T) {
	type testCase[T any] struct {
		name string
		s    *Stack[T]
		want int
	}
	tests := []testCase[bool]{
		{
			name: "Stack with elements size",
			s: &Stack[bool]{
				elements: []bool{true, false, true},
				capacity: 10,
			},
			want: 3,
		},
		{
			name: "Empty stack size",
			s: &Stack[bool]{
				elements: []bool{},
				capacity: 10,
			},
			want: 0,
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
