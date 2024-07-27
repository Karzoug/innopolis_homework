package repo

type CMap[K comparable, V any] interface {
	Load(key K) (value V, ok bool)
	LoadOrStore(key K, value V) (actual V, loaded bool)
	Store(key K, value V)
	Delete(key K)
}

type MapQueues[K comparable, V any] interface {
	LPush(key K, value V)
	LPop(key K) (V, bool)
	LRange(key K, n int) ([]V, bool)
	Keys() []K
	Len() int
}
