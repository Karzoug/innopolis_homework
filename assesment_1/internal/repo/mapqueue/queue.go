package mapqueue

// queue is a unbounded queue.
type queue[T any] struct {
	head *node[T]
	tail *node[T]
}

type node[T any] struct {
	value T
	next  *node[T]
}

// newQueue returns an empty queue.
func newQueue[T any]() *queue[T] {
	n := &node[T]{}
	return &queue[T]{head: n, tail: n}
}

// Enqueue puts the given value v at the tail of the queue.
func (q *queue[T]) Enqueue(v T) {
	n := &node[T]{value: v}
	q.tail.next = n // Link node at the end of the linked list
	q.tail = n      // Swing Tail to node
}

// Dequeue removes and returns the value at the head of the queue.
// It returns false if the queue is empty.
func (q *queue[T]) Dequeue() (T, bool) {
	var t T
	n := q.head
	newHead := n.next
	if newHead == nil {
		return t, false
	}
	v := newHead.value
	newHead.value = t
	q.head = newHead
	return v, true
}

// Range removes and returns the n values at the head of the queue.
// If n equals -1, it returns all values in the queue.
func (q *queue[T]) Range(n int) []T {
	var t T
	arr := make([]T, 0)

	for i := 0; ; i++ {
		h := q.head
		newHead := h.next
		if newHead == nil {
			return arr
		}
		arr = append(arr, newHead.value)
		newHead.value = t
		q.head = newHead

		if n > 0 && i >= n {
			break
		}
	}

	return arr
}
