package queues

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type testCase[T any] struct {
		name string
		want *Queue[T]
	}
	tests := []testCase[string]{
		{
			name: "New queue",
			want: &Queue[string]{elements: []string{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New[string](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Dequeue(t *testing.T) {
	type testCase[T any] struct {
		name    string
		q       Queue[T]
		want    T
		wantErr bool
	}
	tests := []testCase[string]{
		{
			name: "Dequeue item from queue",
			q: Queue[string]{
				elements: []string{"Hello", "World"},
			},
			want:    "Hello",
			wantErr: false,
		},
		{
			name: "Dequeue item from empty queue",
			q: Queue[string]{
				elements: []string{},
			},
			want:    "Hello",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.q.Dequeue()
			if (err != nil) != tt.wantErr {
				t.Errorf("Dequeue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dequeue() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Enqueue(t *testing.T) {
	type args[T any] struct {
		element T
	}
	type testCase[T any] struct {
		name string
		q    Queue[T]
		args args[T]
		want Queue[T]
	}
	tests := []testCase[string]{
		{
			name: "Append element to empty queue",
			q: Queue[string]{
				elements: []string{},
			},
			args: args[string]{"Hello"},
			want: Queue[string]{
				elements: []string{"Hello"},
			},
		},
		{
			name: "Append element to queue",
			q: Queue[string]{
				elements: []string{"Hello"},
			},
			args: args[string]{"World"},
			want: Queue[string]{
				elements: []string{"Hello", "World"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.q.Enqueue(tt.args.element)
		})
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	type testCase[T any] struct {
		name string
		q    Queue[T]
		want bool
	}
	tests := []testCase[string]{
		{
			name: "Queue is empty",
			q: Queue[string]{
				elements: []string{},
			},
			want: true,
		},
		{
			name: "Queue is not empty",
			q: Queue[string]{
				elements: []string{"Hello", "World"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Peek(t *testing.T) {
	type testCase[T any] struct {
		name    string
		q       Queue[T]
		want    T
		wantErr bool
	}
	tests := []testCase[string]{
		{
			name: "Get first element of the queue",
			q: Queue[string]{
				elements: []string{"Hello", "World"},
			},
			want:    "Hello",
			wantErr: false,
		},
		{
			name: "Get error when queue is empty",
			q: Queue[string]{
				elements: []string{},
			},
			want:    "Hello",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.q.Peek()
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

func TestQueue_Size(t *testing.T) {
	type testCase[T any] struct {
		name string
		q    Queue[T]
		want int
	}
	tests := []testCase[string]{
		{
			name: "Empty queue length",
			q: Queue[string]{
				elements: []string{},
			},
			want: 0,
		},
		{
			name: "Queue with elements length",
			q: Queue[string]{
				elements: []string{"Hello", "World"},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
