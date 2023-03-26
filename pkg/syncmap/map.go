package syncmap

import "sync"

type Map struct {
	sync.Mutex
	m map[int]any
}

func NewMap(m map[int]any) *Map {
	return &Map{
		m: m,
	}
}

func (l *Map) Map() any {
	return l.m
}

func (l *Map) Store(key int, value any) {
	l.Lock()
	defer l.Unlock()

	l.m[key] = value
}

func (l *Map) Get(key int) any {
	l.Lock()
	defer l.Unlock()

	if v, has := l.m[key]; has {
		return v
	}
	return nil
}
