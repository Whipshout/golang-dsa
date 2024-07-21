package arrays

import (
	"reflect"
	"testing"
)

func TestArray_Append(t *testing.T) {
	type args[T any] struct {
		element T
	}
	type testCase[T any] struct {
		name string
		a    Array[T]
		args args[T]
		want Array[T]
	}
	tests := []testCase[int]{
		{
			name: "Append to empty array",
			a:    *New[int](),
			args: args[int]{element: 5},
			want: Array[int]{
				data:     []int{5},
				size:     1,
				capacity: 10,
			},
		},
		{
			name: "Append to non-empty array",
			a: Array[int]{
				data:     []int{1, 2, 3},
				size:     3,
				capacity: 10,
			},
			args: args[int]{element: 4},
			want: Array[int]{
				data:     []int{1, 2, 3, 4},
				size:     4,
				capacity: 10,
			},
		},
		{
			name: "Append causing resize",
			a: Array[int]{
				data:     []int{1, 2, 3, 4, 5},
				size:     5,
				capacity: 5,
			},
			args: args[int]{element: 6},
			want: Array[int]{
				data:     []int{1, 2, 3, 4, 5, 6},
				size:     6,
				capacity: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.Append(tt.args.element)
			if !reflect.DeepEqual(tt.a, tt.want) {
				t.Errorf("After Append() = %v, want %v", tt.a, tt.want)
			}
		})
	}
}

func TestArray_Get(t *testing.T) {
	type args struct {
		index int
	}
	type testCase[T any] struct {
		name    string
		a       Array[T]
		args    args
		want    T
		wantErr bool
	}
	tests := []testCase[int]{
		{
			name: "Get item from array",
			a: Array[int]{
				data:     []int{1, 2, 3},
				size:     3,
				capacity: 10,
			},
			args:    args{index: 0},
			want:    1,
			wantErr: false,
		},
		{
			name: "Error getting item from array",
			a: Array[int]{
				data:     []int{1, 2, 3, 4, 5, 6},
				size:     6,
				capacity: 10,
			},
			args:    args{index: 8},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.Get(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArray_Set(t *testing.T) {
	type args[T any] struct {
		index   int
		element T
	}
	type testCase[T any] struct {
		name    string
		a       Array[T]
		args    args[T]
		want    Array[T]
		wantErr bool
	}
	tests := []testCase[int]{
		{
			name: "Set item from array",
			a: Array[int]{
				data:     []int{1, 2, 3},
				size:     3,
				capacity: 10,
			},
			args: args[int]{index: 0, element: 5},
			want: Array[int]{
				data:     []int{5, 2, 3},
				size:     3,
				capacity: 10,
			},
			wantErr: false,
		},
		{
			name: "Error setting item from array",
			a: Array[int]{
				data:     []int{1, 2, 3, 4, 5, 6},
				size:     6,
				capacity: 10,
			},
			args: args[int]{index: 8, element: 10},
			want: Array[int]{
				data:     []int{1, 2, 3, 4, 5, 6},
				size:     6,
				capacity: 10,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.a.Set(tt.args.index, tt.args.element)
			if (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.a, tt.want) {
				t.Errorf("After Set() = %v, want %v", tt.a, tt.want)
			}
		})
	}
}

func TestArray_Size(t *testing.T) {
	type testCase[T any] struct {
		name string
		a    Array[T]
		want int
	}
	tests := []testCase[int]{
		{
			name: "Get item from array",
			a: Array[int]{
				data:     []int{1, 2, 3},
				size:     3,
				capacity: 10,
			},
			want: 3,
		},
		{
			name: "Get error getting item from array",
			a: Array[int]{
				data:     []int{1, 2, 3, 4, 5, 6},
				size:     6,
				capacity: 10,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArray_resize(t *testing.T) {
	type testCase[T any] struct {
		name     string
		a        Array[T]
		wantCap  int
		wantSize int
	}
	tests := []testCase[int]{
		{
			name: "Resize empty array",
			a: Array[int]{
				data:     make([]int, 0, 10),
				size:     0,
				capacity: 10,
			},
			wantCap:  20,
			wantSize: 0,
		},
		{
			name: "Resize partially filled array",
			a: Array[int]{
				data:     []int{1, 2, 3, 4, 5},
				size:     5,
				capacity: 10,
			},
			wantCap:  20,
			wantSize: 5,
		},
		{
			name: "Resize full array",
			a: Array[int]{
				data:     []int{1, 2, 3, 4, 5},
				size:     5,
				capacity: 5,
			},
			wantCap:  10,
			wantSize: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.resize()
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		initialCapacity []int
	}
	type testCase[T any] struct {
		name string
		args args
		want *Array[T]
	}
	tests := []testCase[int]{
		{
			name: "New array with default capacity",
			args: args{initialCapacity: []int{}},
			want: &Array[int]{
				data:     make([]int, 0, 10),
				size:     0,
				capacity: 10,
			},
		},
		{
			name: "New array with custom capacity",
			args: args{initialCapacity: []int{20}},
			want: &Array[int]{
				data:     make([]int, 0, 20),
				size:     0,
				capacity: 20,
			},
		},
		{
			name: "New array with zero capacity (should use default)",
			args: args{initialCapacity: []int{0}},
			want: &Array[int]{
				data:     make([]int, 0, 10),
				size:     0,
				capacity: 10,
			},
		},
		{
			name: "New array with negative capacity (should use default)",
			args: args{initialCapacity: []int{-5}},
			want: &Array[int]{
				data:     make([]int, 0, 10),
				size:     0,
				capacity: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New[int](tt.args.initialCapacity...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArray_Pop(t *testing.T) {
	type testCase[T any] struct {
		name    string
		a       Array[T]
		want    T
		wantErr bool
		wantLen int
	}
	tests := []testCase[int]{
		{
			name: "Pop from non-empty array",
			a: Array[int]{
				data:     []int{1, 2, 3, 4, 5},
				size:     5,
				capacity: 10,
			},
			want:    5,
			wantErr: false,
			wantLen: 4,
		},
		{
			name: "Pop from array with one element",
			a: Array[int]{
				data:     []int{1},
				size:     1,
				capacity: 10,
			},
			want:    1,
			wantErr: false,
			wantLen: 0,
		},
		{
			name: "Pop from empty array",
			a: Array[int]{
				data:     []int{},
				size:     0,
				capacity: 10,
			},
			want:    0, // zero value for int
			wantErr: true,
			wantLen: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() got = %v, want %v", got, tt.want)
			}
			if tt.a.size != tt.wantLen {
				t.Errorf("After Pop(): size = %v, want %v", tt.a.size, tt.wantLen)
			}
			if len(tt.a.data) != tt.wantLen {
				t.Errorf("After Pop(): len(data) = %v, want %v", len(tt.a.data), tt.wantLen)
			}
		})
	}
}
