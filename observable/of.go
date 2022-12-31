package observable

import "github.com/lyyourc/straw"

func Of[T any](values ...T) straw.Observable[T] {
	return straw.NewObservable(func(s straw.Subscriber[T]) straw.Subscription {
		for _, v := range values {
			s.Next(v)
		}

		s.Complete()

		return straw.Subscription{}
	})
}
