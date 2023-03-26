package syncmap

import "sync"

type Map[K comparable, V any] struct {
	sync.Mutex
	m map[K]V
}

func NewMap[K comparable, V any](m map[K]V) *Map[K, V] {
	return &Map[K, V]{
		m: m,
	}
}

func (l *Map[K, V]) Map() map[K]V {
	return l.m
}

func (l *Map[K, V]) Store(key K, value V) {
	l.Lock()
	defer l.Unlock()

	l.m[key] = value
}

func (l *Map[K, V]) Get(key K) V {
	l.Lock()
	defer l.Unlock()

	if v, has := l.m[key]; has {
		return v
	}

	// workaround to zero value of V
	var n V
	return n
}
