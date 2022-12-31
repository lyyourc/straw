package operator

import "github.com/lyyourc/straw"

type SwitchMapFunc[T any, R any] func(val T, index int) straw.Observable[R]

func SwitchMap[T any, R any](project SwitchMapFunc[T, R]) straw.OperatorFunc[T, R] {
	return func(source straw.Observable[T]) straw.Observable[R] {
		return straw.NewObservable(func(o straw.Subscriber[R]) straw.Subscription {
			i := 0

			source.Subscribe(straw.Observer[T]{
				Next: func(value T) {
					innerObservable := project(value, i)
					i += 1
					innerObservable.Subscribe(straw.Observer[R]{
						Next: func(val R) {
							o.Next(val)
						},
						Error: func(err error) {
							o.Error(err)
						},
						Complete: func() {
							o.Complete()
						},
					})
				},
				Error: func(err error) {
					// dont care
				},
			})

			return straw.Subscription{}
		})
	}
}
