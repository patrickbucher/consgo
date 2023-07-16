package consgo

type Func[T any] func(left T) T

type Cell[T any] struct {
	left  T
	right *Cell[T]
}

func Cons[T any](left T, right *Cell[T]) *Cell[T] {
	return &Cell[T]{
		left:  left,
		right: right,
	}
}

func Car[T any](cell *Cell[T]) T {
	return cell.left
}

func Cdr[T any](cell *Cell[T]) *Cell[T] {
	return cell.right
}

func Ref[T any](cell *Cell[T], n int) T {
	if n == 0 {
		return Car(cell)
	}
	return Ref(Cdr(cell), n-1)
}

func Map[T any](cell *Cell[T], f Func[T]) *Cell[T] {
	if cell == nil {
		return nil
	}
	return Cons(f(Car(cell)), Map(Cdr(cell), f))
}

func ForEach[T any](cell *Cell[T], f func(x T)) {
	if cell == nil {
		return
	}
	f(Car(cell))
	ForEach(Cdr(cell), f)
}

func Slice[T any](head *Cell[T]) []T {
	slice := make([]T, 0)
	for head != nil {
		slice = append(slice, Car(head))
		head = Cdr(head)
	}
	return slice
}

func List[T any](slice []T) *Cell[T] {
	var head *Cell[T]
	n := len(slice)
	if n == 0 {
		return nil
	}
	for i := n - 1; i >= 0; i-- {
		head = Cons(slice[i], head)
	}
	return head
}
