package linked_lists

import (
	"reflect"
	"testing"
)

func TestLinkedList_Back(t *testing.T) {
	type testCase[T any] struct {
		name    string
		ll      *LinkedList[T]
		want    T
		wantErr bool
	}
	linkedListWithItems := New[int]()
	linkedListWithItems.PushFront(1)
	linkedListWithItems.PushBack(2)
	tests := []testCase[int]{
		{
			name:    "Back item of a linked list with items",
			ll:      linkedListWithItems,
			want:    2,
			wantErr: false,
		},
		{
			name:    "Back item of an empty linked list",
			ll:      New[int](),
			want:    1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ll.Back()
			if (err != nil) != tt.wantErr {
				t.Errorf("Back() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Back() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Clear(t *testing.T) {
	type testCase[T any] struct {
		name string
		ll   *LinkedList[T]
	}
	linkedListWithItems := New[int]()
	linkedListWithItems.PushFront(1)
	linkedListWithItems.PushBack(2)
	tests := []testCase[int]{
		{
			name: "Clear a linked list",
			ll:   linkedListWithItems,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ll.Clear()
			if !reflect.DeepEqual(New[int](), tt.ll) {
				t.Errorf("got %v, want %v", New[int](), tt.ll)
			}
		})
	}
}

func TestLinkedList_Front(t *testing.T) {
	type testCase[T any] struct {
		name    string
		ll      *LinkedList[T]
		want    T
		wantErr bool
	}
	linkedListWithItems := New[int]()
	linkedListWithItems.PushFront(1)
	linkedListWithItems.PushBack(2)
	tests := []testCase[int]{
		{
			name:    "Front item of a linked list with items",
			ll:      linkedListWithItems,
			want:    1,
			wantErr: false,
		},
		{
			name:    "Front item of an empty linked list",
			ll:      New[int](),
			want:    1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ll.Front()
			if (err != nil) != tt.wantErr {
				t.Errorf("Front() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Front() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_IsEmpty(t *testing.T) {
	type testCase[T any] struct {
		name string
		ll   *LinkedList[T]
		want bool
	}
	linkedListWithItems := New[int]()
	linkedListWithItems.PushFront(1)
	tests := []testCase[int]{
		{
			name: "Linked list is empty",
			ll:   New[int](),
			want: true,
		},
		{
			name: "Linked list is not empty",
			ll:   linkedListWithItems,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ll.IsEmpty(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_PopBack(t *testing.T) {
	type testCase[T any] struct {
		name    string
		ll      *LinkedList[T]
		want    T
		wantErr bool
	}
	linkedListWithItems := New[int]()
	linkedListWithItems.PushFront(1)
	linkedListWithItems.PushBack(2)
	tests := []testCase[int]{
		{
			name:    "PopBack from linked list with items",
			ll:      linkedListWithItems,
			want:    2,
			wantErr: false,
		},
		{
			name:    "PopBack from empty linked list",
			ll:      New[int](),
			want:    1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ll.PopBack()
			if (err != nil) != tt.wantErr {
				t.Errorf("PopBack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopBack() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_PopFront(t *testing.T) {
	type testCase[T any] struct {
		name    string
		ll      *LinkedList[T]
		want    T
		wantErr bool
	}
	linkedListWithItems := New[int]()
	linkedListWithItems.PushFront(1)
	linkedListWithItems.PushBack(2)
	tests := []testCase[int]{
		{
			name:    "PopFront from linked list with items",
			ll:      linkedListWithItems,
			want:    1,
			wantErr: false,
		},
		{
			name:    "PopFront from empty linked list",
			ll:      New[int](),
			want:    1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ll.PopFront()
			if (err != nil) != tt.wantErr {
				t.Errorf("PopFront() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopFront() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_PushBack(t *testing.T) {
	type args[T any] struct {
		value T
	}
	linkedListWithItems := New[int]()
	linkedListWithItems.PushFront(1)
	linkedListWithItems.PushBack(2)
	type testCase[T any] struct {
		name       string
		ll         *LinkedList[T]
		args       args[T]
		backWanted int
	}
	tests := []testCase[int]{
		{
			name:       "PushBack item to a not empty linked list",
			ll:         linkedListWithItems,
			args:       args[int]{value: 3},
			backWanted: 3,
		},
		{
			name:       "PushBack item to an empty linked list",
			ll:         New[int](),
			args:       args[int]{value: 1},
			backWanted: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ll.PushBack(tt.args.value)
			got, _ := tt.ll.Back()
			if !reflect.DeepEqual(got, tt.backWanted) {
				t.Errorf("LinkedList.PushBack() got = %v, want %v", got, tt.backWanted)
			}
		})
	}
}

func TestLinkedList_PushFront(t *testing.T) {
	type args[T any] struct {
		value T
	}
	linkedListWithItems := New[int]()
	linkedListWithItems.PushFront(1)
	linkedListWithItems.PushBack(2)
	type testCase[T any] struct {
		name        string
		ll          *LinkedList[T]
		args        args[T]
		frontWanted int
	}
	tests := []testCase[int]{
		{
			name:        "PushFront item to a not empty linked list",
			ll:          linkedListWithItems,
			args:        args[int]{value: 3},
			frontWanted: 3,
		},
		{
			name:        "PushFront item to an empty linked list",
			ll:          New[int](),
			args:        args[int]{value: 1},
			frontWanted: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ll.PushFront(tt.args.value)
			got, _ := tt.ll.Front()
			if !reflect.DeepEqual(got, tt.frontWanted) {
				t.Errorf("LinkedList.PushFront() got = %v, want %v", got, tt.frontWanted)
			}
		})
	}
}

func TestLinkedList_Size(t *testing.T) {
	type testCase[T any] struct {
		name string
		ll   *LinkedList[T]
		want int
	}
	linkedListWithItems := New[int]()
	linkedListWithItems.PushFront(1)
	linkedListWithItems.PushBack(2)
	tests := []testCase[int]{
		{
			name: "Linked list with items size",
			ll:   linkedListWithItems,
			want: 2,
		},
		{
			name: "Linked list without items size",
			ll:   New[int](),
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ll.Size(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type testCase[T any] struct {
		name string
		want *LinkedList[T]
	}
	tests := []testCase[int]{
		{
			name: "New linked list",
			want: &LinkedList[int]{},
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
