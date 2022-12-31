package observable

import "github.com/lyyourc/straw"

func Zip[T any](observables ...straw.Observable[T]) straw.Observable[[]T] {
	buffers := make([][]T, 0)
	for range observables {
		buffer := make([]T, 0)
		buffers = append(buffers, buffer)
	}

	return straw.NewObservable(func(sub straw.Subscriber[[]T]) straw.Subscription {
		for i, observable := range observables {
			observable.Subscribe(straw.Observer[T]{
				Next: func(value T) {
					buffers[i] = append(buffers[i], value)

					ifEveryHasOneValue := true
					for _, b := range buffers {
						if len(b) == 0 {
							ifEveryHasOneValue = false
							break
						}
					}

					if ifEveryHasOneValue {
						next := make([]T, 0)
						for _, b := range buffers {
							next = append(next, b[0])
							b = b[1:]
						}
						sub.Next(next)
					}
				},
			})
		}

		return straw.Subscription{}
	})
}
