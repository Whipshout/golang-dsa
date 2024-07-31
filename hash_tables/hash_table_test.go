package hash_tables

import (
	"reflect"
	"testing"
)

func TestHashTable_Delete(t *testing.T) {
	type args[K comparable] struct {
		key K
	}
	type testCase[K comparable, V any] struct {
		name string
		ht   *HashTable[K, V]
		args args[K]
		want bool
	}
	tests := []testCase[string, string]{
		{
			name: "Delete existing key",
			ht: func() *HashTable[string, string] {
				ht := New[string, string]()
				ht.Insert("key1", "value1")
				ht.Insert("key2", "value2")
				return ht
			}(),
			args: args[string]{key: "key1"},
			want: true,
		},
		{
			name: "Delete non-existing key",
			ht: func() *HashTable[string, string] {
				ht := New[string, string]()
				ht.Insert("key1", "value1")
				return ht
			}(),
			args: args[string]{key: "key2"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ht.Delete(tt.args.key); got != tt.want {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashTable_Display(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name string
		ht   *HashTable[K, V]
	}
	tests := []testCase[string, string]{
		{
			name: "Display hash table",
			ht: func() *HashTable[string, string] {
				ht := New[string, string]()
				ht.Insert("key1", "value1")
				ht.Insert("key2", "value2")
				return ht
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ht.Display()
		})
	}
}

func TestHashTable_Get(t *testing.T) {
	type args[K comparable] struct {
		key K
	}
	type testCase[K comparable, V any] struct {
		name  string
		ht    *HashTable[K, V]
		args  args[K]
		want  V
		want1 bool
	}
	tests := []testCase[string, string]{
		{
			name: "Get existing key",
			ht: func() *HashTable[string, string] {
				ht := New[string, string]()
				ht.Insert("key1", "value1")
				return ht
			}(),
			args:  args[string]{key: "key1"},
			want:  "value1",
			want1: true,
		},
		{
			name: "Get non-existing key",
			ht: func() *HashTable[string, string] {
				ht := New[string, string]()
				ht.Insert("key1", "value1")
				return ht
			}(),
			args:  args[string]{key: "key2"},
			want:  "",
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.ht.Get(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestHashTable_Insert(t *testing.T) {
	type args[K comparable, V any] struct {
		key   K
		value V
	}
	type testCase[K comparable, V any] struct {
		name string
		ht   *HashTable[K, V]
		args args[K, V]
	}
	tests := []testCase[string, string]{
		{
			name: "Insert new key",
			ht:   New[string, string](),
			args: args[string, string]{key: "key1", value: "value1"},
		},
		{
			name: "Insert existing key with new value",
			ht: func() *HashTable[string, string] {
				ht := New[string, string]()
				ht.Insert("key1", "value1")
				return ht
			}(),
			args: args[string, string]{key: "key1", value: "value2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ht.Insert(tt.args.key, tt.args.value)
		})
	}
}

func TestHashTable_Size(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name string
		ht   *HashTable[K, V]
		want int
	}
	tests := []testCase[string, string]{
		{
			name: "Size after inserts",
			ht: func() *HashTable[string, string] {
				ht := New[string, string]()
				ht.Insert("key1", "value1")
				ht.Insert("key2", "value2")
				return ht
			}(),
			want: 2,
		},
		{
			name: "Size after delete",
			ht: func() *HashTable[string, string] {
				ht := New[string, string]()
				ht.Insert("key1", "value1")
				ht.Insert("key2", "value2")
				ht.Delete("key1")
				return ht
			}(),
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ht.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHashTable(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name string
		want *HashTable[K, V]
	}
	tests := []testCase[string, string]{
		{
			name: "New hash table",
			want: New[string, string](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New[string, string](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHashTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hash(t *testing.T) {
	type args[K comparable] struct {
		key K
	}
	type testCase[K comparable] struct {
		name string
		args args[K]
		want int
	}
	tests := []testCase[string]{
		{
			name: "Hash string key",
			args: args[string]{key: "key1"},
			want: hash("key1"),
		},
		{
			name: "Hash int key",
			args: args[string]{key: "42"},
			want: hash("42"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hash(tt.args.key); got != tt.want {
				t.Errorf("hash() = %v, want %v", got, tt.want)
			}
		})
	}
}
