package mxmap

import (
	"sync"

	"assesment_1/internal/repo"
)

var _ repo.CMap[int, int] = (*MxMap[int, int])(nil)

type MxMap[K comparable, V any] struct {
	wl map[K]V
	mx sync.Mutex
}

func New[K comparable, V any]() *MxMap[K, V] {
	return &MxMap[K, V]{
		wl: make(map[K]V),
		mx: sync.Mutex{},
	}
}

func (m *MxMap[K, V]) Load(token K) (V, bool) {
	m.mx.Lock()
	defer m.mx.Unlock()

	v, ok := m.wl[token]
	return v, ok
}

func (m *MxMap[K, V]) Store(token K, value V) {
	m.mx.Lock()
	defer m.mx.Unlock()

	m.wl[token] = value
}

func (m *MxMap[K, V]) LoadOrStore(token K, value V) (V, bool) {
	m.mx.Lock()
	defer m.mx.Unlock()

	v, ok := m.wl[token]
	if ok {
		return v, true
	}

	m.wl[token] = value
	return value, false
}

func (m *MxMap[K, V]) Delete(token K) {
	m.mx.Lock()
	defer m.mx.Unlock()

	delete(m.wl, token)
}
