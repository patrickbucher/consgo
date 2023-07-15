package consgo

type Func[T any] func(left T) T
type Delayed[T any] func() *Stream[T]

type Stream[T any] struct {
	left  T
	right Delayed[T]
}

func New[T any](left T, next Func[T]) *Stream[T] {
	right := func() *Stream[T] { return New(next(left), next) }
	// TODO: memoize right
	return &Stream[T]{
		left:  left,
		right: right,
	}
}

func ConsStream[T any](left T, right Delayed[T]) *Stream[T] {
	return &Stream[T]{
		left:  left,
		right: right,
	}
}

func StreamCar[T any](stream *Stream[T]) T {
	return stream.left
}

func StreamCdr[T any](stream *Stream[T]) *Stream[T] {
	return stream.right()
}

func StreamRef[T any](stream *Stream[T], n int) T {
	if n == 0 {
		return StreamCar(stream)
	}
	return StreamRef(StreamCdr(stream), n-1)
}

func StreamTake[T any](s *Stream[T], n int) []T {
	values := make([]T, 0)
	for i := 0; i < n; i++ {
		values = append(values, StreamCar(s))
		s = StreamCdr(s)
	}
	return values
}
