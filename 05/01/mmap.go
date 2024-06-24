package main

import (
	"fmt"
	"sync"
)

type mmap[K comparable, V any] struct {
	mu   sync.Mutex
	data map[K]V
}

func NewMmap[K comparable, V any](data map[K]V) *mmap[K, V] {
	if data == nil {
		data = make(map[K]V)
	}
	return &mmap[K, V]{
		data: data,
		mu:   sync.Mutex{},
	}
}

func (m *mmap[K, V]) Map(f func(key K, value V) V) {
	m.mu.Lock()
	for key, value := range m.data {
		m.data[key] = f(key, value)
	}
	m.mu.Unlock()
}

func (m *mmap[K, V]) Set(key K, value V) {
	m.mu.Lock()
	m.data[key] = value
	m.mu.Unlock()
}
func (m *mmap[K, V]) Get(key K) (V, bool) {
	m.mu.Lock()
	k, ok := m.data[key]
	m.mu.Unlock()

	return k, ok
}

func (m *mmap[K, V]) Delete(key K) {
	m.mu.Lock()
	delete(m.data, key)
	m.mu.Unlock()
}

func (m *mmap[K, V]) Clear() {
	m.mu.Lock()
	clear(m.data)
	m.mu.Unlock()
}

func (m *mmap[K, V]) String() string {
	m.mu.Lock()
	s := fmt.Sprint(m.data)
	m.mu.Unlock()

	return s
}
