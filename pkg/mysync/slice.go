package mysync

import "sync"

type Slice[V any] struct {
	sync.Mutex
	s []V
}

func NewSlice[V any](s []V) *Slice[V] {
	return &Slice[V]{
		s: s,
	}
}

func (l *Slice[V]) Slice() []V {
	return l.s
}

func (l *Slice[V]) Append(value V) {
	l.Lock()
	defer l.Unlock()

	l.s = append(l.s, value)
}

func (l *Slice[V]) Get(index int) V {
	l.Lock()
	defer l.Unlock()

	return l.s[index]
}
