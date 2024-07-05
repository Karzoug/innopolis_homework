package mapqueue

import (
	"sync"

	"assesment_1/internal/repo"

	"golang.org/x/exp/maps"
)

var _ repo.MapQueues[string, any] = (*mapQueues[string, any])(nil)

// mapQueues is a concurrent map of lists. Methods are thread-safe.
type mapQueues[K comparable, V any] struct {
	m  map[K]*queue[V]
	mx sync.Mutex
}

// New creates a new MapQueues instance.
func New[K comparable, V any]() *mapQueues[K, V] {
	return &mapQueues[K, V]{
		m:  make(map[K]*queue[V]),
		mx: sync.Mutex{},
	}
}

// LPush inserts the specified value at the head of the queue stored at key.
func (c *mapQueues[K, V]) LPush(key K, value V) {
	var (
		q  *queue[V]
		ok bool
	)

	c.mx.Lock()
	defer c.mx.Unlock()

	q, ok = c.m[key]
	if !ok {
		q = newQueue[V]()
		c.m[key] = q
	}
	//c.count += len(values)
	q.Enqueue(value)
}

// LPop removes and returns the first elements of the queue stored at key.
func (c *mapQueues[K, V]) LPop(key K) (V, bool) {
	var v V

	c.mx.Lock()
	defer c.mx.Unlock()

	q, ok := c.m[key]
	if !ok {
		return v, false
	}

	return q.Dequeue()
}

func (c *mapQueues[K, V]) Keys() []K {
	c.mx.Lock()
	defer c.mx.Unlock()

	return maps.Keys(c.m)
}

// Range returns the specified elements of the list stored at key.
func (c *mapQueues[K, V]) LRange(key K, n int) ([]V, bool) {
	c.mx.Lock()
	defer c.mx.Unlock()

	q, ok := c.m[key]
	if !ok {
		return []V{}, false
	}

	v := q.Range(n)
	if len(v) == 0 {
		delete(c.m, key)
		return v, false
	}

	return v, true
}

// Len returns the number of queues in the map.
func (c *mapQueues[K, V]) Len() int {
	c.mx.Lock()
	defer c.mx.Unlock()

	return len(c.m)
}
