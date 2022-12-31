package operator

import "github.com/lyyourc/straw"

type FilterFunc[T any] func(val T, index int) bool

func Filter[T any](filter FilterFunc[T]) straw.OperatorFunc[T, T] {
	return func(source straw.Observable[T]) straw.Observable[T] {
		innerObservable := straw.NewObservable(func(o straw.Subscriber[T]) straw.Subscription {
			i := 0

			source.Subscribe(straw.Observer[T]{
				Next: func(val T) {
					result := filter(val, i)
					if result {
						i += 1
						o.Next(val)
					}
				},
				Error: func(err error) {
					o.Error(err)
				},
				Complete: func() {
					o.Complete()
				},
			})

			return straw.Subscription{}
		})

		return innerObservable
	}
}
