// A sorted set with a fixed size (smallest --> biggest).
package fixedset;

import (
	"sort"
)

type FixedSet[T any] struct {
	items []T
	less func(a, b T) bool;
}

func New[T any](maxSize int, less func(a, b T) bool) *FixedSet[T] {
	return &FixedSet[T]{make([]T, 0, maxSize+1), less};
}

func (f *FixedSet[T]) Len() int {
	return len(f.items)
}

func (f *FixedSet[T]) Iterate() []T {
	return f.items;
}

func (f *FixedSet[T]) Add(t T) bool {
	l := len(f.items);
	if l == 0 {
		f.items = append(f.items, t);
		return true;
	}
	if l < cap(f.items) || f.less(t, f.items[l-1]) {
		i := sort.Search(l, func(i int) bool {
			return !f.less(f.items[i], t);
		});
		f.items = append(f.items, t);
		copy(f.items[i+1:], f.items[i:]);
		f.items[i] = t;
		if len(f.items) == cap(f.items) {
			f.items = f.items[0:cap(f.items)-1];
		}
		return true;
	}
	return false;
}
