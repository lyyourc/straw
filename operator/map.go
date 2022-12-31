package operator

import "github.com/lyyourc/straw"

type MapFunc[T any, R any] func(value T, index int) R

func Map[T any, R any](project MapFunc[T, R]) straw.OperatorFunc[T, R] {
	return func(source straw.Observable[T]) straw.Observable[R] {
		return straw.NewObservable[R](func(s straw.Subscriber[R]) straw.Subscription {
			i := 0

			source.Subscribe(straw.Observer[T]{
				Next: func(value T) {
					nextValue := project(value, i)
					i += 1
					s.Next(nextValue)
				},
				Error: func(err error) {
					s.Error(err)
				},
				Complete: func() {
					s.Complete()
				},
			})

			return straw.Subscription{}
		})

	}
}
