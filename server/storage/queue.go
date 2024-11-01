package storage

type Queue[T any] chan T

func (self Queue[T]) Count() int {
	return len(self)
}

func (self *Queue[T]) Push(item T) {
	select {
	case *self <- item:
	default:
		queue := make(chan T, len(*self)*2)

		for len(*self) > 0 {
			queue <- <-*self
		}

		*self = queue
		*self <- item
	}
}
