package stream

type StreamInterface[T any] interface {
	Listen(callback func(T))
	Filter(callback func(T) bool) Stream[T]
	Reduce(callback func(T, T) T) T
	ToArray() []T
}

type Stream[T any] struct {
	data []T
}

func NewStream[T any](data []T) Stream[T] {
	return Stream[T]{data: data}
}

func NewStreamEmpty[T any]() Stream[T] {
	return Stream[T]{data: []T{}}
}

func (s *Stream[T]) Listen(callback func(T)) {
	for _, item := range s.data {
		callback(item)
	}
}

func (s *Stream[T]) Filter(callback func(T) bool) Stream[T] {
	var filtered []T
	for _, item := range s.data {
		if callback(item) {
			filtered = append(filtered, item)
		}
	}
	return Stream[T]{data: filtered}
}

func (s *Stream[T]) Reduce(callback func(T, T) T) T {
	var result T
	for _, item := range s.data {
		result = callback(result, item)
	}
	return result
}

func (s *Stream[T]) ToArray() []T {
	return s.data
}

func (s *Stream[T]) Add(item T) {
	s.data = append(s.data, item)
}
