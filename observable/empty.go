package observable

import "github.com/lyyourc/straw"

func Empty() straw.Observable[string] {
	return straw.NewObservable(func(s straw.Subscriber[string]) straw.Subscription {
		s.Complete()
		return straw.Subscription{}
	})
}
