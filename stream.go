package consgo

type Delayed[T any] func() *Stream[T]

type Stream[T any] struct {
	left  T
	right Delayed[T]
}

func New[T any](left T, next Func[T]) *Stream[T] {
	right := func() *Stream[T] { return New(next(left), next) }
	return &Stream[T]{
		left:  left,
		right: memoize(right),
	}
}

func ConsStream[T any](left T, right Delayed[T]) *Stream[T] {
	return &Stream[T]{
		left:  left,
		right: memoize(right),
	}
}

func StreamCar[T any](stream *Stream[T]) T {
	return stream.left
}

func StreamCdr[T any](stream *Stream[T]) Delayed[T] {
	return stream.right
}

func StreamRef[T any](stream *Stream[T], n int) T {
	if n == 0 {
		return StreamCar(stream)
	}
	return StreamRef(StreamCdr(stream)(), n-1)
}

func StreamMap[T any](stream *Stream[T], f Func[T]) *Stream[T] {
	if stream == nil {
		return nil
	}
	return ConsStream(
		f(StreamCar(stream)),
		func() *Stream[T] {
			return StreamMap(StreamCdr(stream)(), f)
		})
}

// TODO: StreamForEach

func StreamTake[T any](s *Stream[T], n int) []T {
	values := make([]T, 0)
	for i := 0; i < n; i++ {
		values = append(values, StreamCar(s))
		s = StreamCdr(s)()
	}
	return values
}

func memoize[T any](f Delayed[T]) Delayed[T] {
	run := false
	var result *Stream[T]
	return func() *Stream[T] {
		if !run {
			result = f()
			run = true
		}
		return result
	}
}
