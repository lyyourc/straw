package observable

import "github.com/lyyourc/straw"

func Pipe[T any, R any](source straw.Observable[T], fn straw.OperatorFunc[T, R], fns ...straw.OperatorFunc[R, R]) straw.Observable[R] {
	current := fn(source)

	for _, f := range fns {
		current = f(current)
	}

	return current
}
