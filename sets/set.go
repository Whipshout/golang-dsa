package sets

type Set[T comparable] map[T]struct{}

func New[T comparable]() *Set[T] {
	set := make(Set[T])
	return &set
}

func (s Set[T]) Add(element T) {
	s[element] = struct{}{}
}

func (s Set[T]) Remove(element T) {
	delete(s, element)
}

func (s Set[T]) Contains(element T) bool {
	_, exists := s[element]

	return exists
}

func (s Set[T]) Size() int {
	return len(s)
}

func (s Set[T]) Clear() {
	for element := range s {
		delete(s, element)
	}
}

func (s Set[T]) ToSlice() []T {
	slice := make([]T, 0, s.Size())

	for item := range s {
		slice = append(slice, item)
	}

	return slice
}

func (s Set[T]) Union(other Set[T]) *Set[T] {
	unionSet := New[T]()

	for element := range s {
		unionSet.Add(element)
	}
	for element := range other {
		unionSet.Add(element)
	}

	return unionSet
}

func (s Set[T]) Intersection(other Set[T]) *Set[T] {
	intersectionSet := New[T]()

	for element := range s {
		if other.Contains(element) {
			intersectionSet.Add(element)
		}
	}

	return intersectionSet
}

func (s Set[T]) Difference(other Set[T]) *Set[T] {
	differenceSet := New[T]()

	for element := range s {
		if !other.Contains(element) {
			differenceSet.Add(element)
		}
	}

	return differenceSet
}
