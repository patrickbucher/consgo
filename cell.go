package consgo

type Cell[T any] struct {
	Left  T
	Right *Cell[T]
}

func Cons[T any](left T, right *Cell[T]) *Cell[T] {
	return &Cell[T]{
		Left:  left,
		Right: right,
	}
}

func Car[T any](cell *Cell[T]) T {
	return cell.Left
}

func Cdr[T any](cell *Cell[T]) *Cell[T] {
	return cell.Right
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
