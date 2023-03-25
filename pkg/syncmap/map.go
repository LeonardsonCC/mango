package syncmap

import "sync"

type Map struct {
	sync.Mutex
	m map[int][]byte
}

func NewMap(m map[int][]byte) *Map {
	return &Map{
		m: m,
	}
}

func (l *Map) Map() map[int][]byte {
	return l.m
}

func (l *Map) Store(key int, value []byte) {
	l.Lock()
	defer l.Unlock()

	l.m[key] = value
}

func (l *Map) Get(key int) []byte {
	l.Lock()
	defer l.Unlock()

	if v, has := l.m[key]; has {
		return v
	}
	return nil
}
