package heaps

import (
	"golang.org/x/exp/constraints"
	"reflect"
	"testing"
)

func TestHeap_Pop(t *testing.T) {
	type testCase[T constraints.Ordered] struct {
		name     string
		h        *Heap[T]
		want     T
		wantErr  bool
		wantHeap *Heap[T]
	}
	heapWithElements := New[int]()
	heapWithElements.Push(1)
	heapWithElements.Push(3)
	heapWithElements.Push(2)
	heapWithElements.Push(4)
	sortedHeapWanted := &Heap[int]{elements: []int{1, 2, 3}}
	tests := []testCase[int]{
		{
			name:     "Pop to an empty heap",
			h:        New[int](),
			want:     0,
			wantErr:  true,
			wantHeap: nil,
		},
		{
			name:     "Pop to a heap",
			h:        heapWithElements,
			want:     4,
			wantErr:  false,
			wantHeap: sortedHeapWanted,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.h.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeap_Push(t *testing.T) {
	type args[T constraints.Ordered] struct {
		element T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		h    *Heap[T]
		args args[T]
		want *Heap[T]
	}
	heapWithElements := New[int]()
	heapWithElements.Push(1)
	heapWithElements.Push(3)
	heapWithElements.Push(2)
	sortedHeapWanted := &Heap[int]{elements: []int{1, 2, 3, 4}}
	tests := []testCase[int]{
		{
			name: "Push to an empty heap",
			h:    New[int](),
			args: args[int]{element: 1},
			want: &Heap[int]{elements: []int{1}},
		},
		{
			name: "Push to a heap",
			h:    heapWithElements,
			args: args[int]{element: 4},
			want: &Heap[int]{elements: []int{1, 2, 3, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Push(tt.args.element)
			if tt.h.Size() > 1 && !reflect.DeepEqual(sortedHeapWanted, tt.want) {
				t.Errorf("Push() = %v, want %v", sortedHeapWanted, tt.want)
			}
		})
	}
}

func TestHeap_Size(t *testing.T) {
	type testCase[T constraints.Ordered] struct {
		name string
		h    *Heap[T]
		want int
	}
	heapWithElements := New[int]()
	heapWithElements.Push(1)
	heapWithElements.Push(2)
	tests := []testCase[int]{
		{
			name: "Size of an empty heap",
			h:    New[int](),
			want: 0,
		},
		{
			name: "Size of a heap",
			h:    heapWithElements,
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type testCase[T constraints.Ordered] struct {
		name string
		want *Heap[T]
	}
	tests := []testCase[int]{
		{
			name: "New heap",
			want: &Heap[int]{elements: make([]int, 0)},
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

func TestHeap_Peek(t *testing.T) {
	type testCase[T constraints.Ordered] struct {
		name    string
		h       *Heap[T]
		want    T
		wantErr bool
	}
	heapWithElements := New[int]()
	heapWithElements.Push(1)
	heapWithElements.Push(2)
	tests := []testCase[int]{
		{
			name:    "Peak of an empty heap",
			h:       New[int](),
			want:    0,
			wantErr: true,
		},
		{
			name:    "Peak of a heap",
			h:       heapWithElements,
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.h.Peek()
			if (err != nil) != tt.wantErr {
				t.Errorf("Peek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Peek() got = %v, want %v", got, tt.want)
			}
		})
	}
}
