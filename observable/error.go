package observable

import "github.com/lyyourc/straw"

func Error(err error) straw.Observable[string] {
	return straw.NewObservable(func(s straw.Subscriber[string]) straw.Subscription {
		s.Error(err)
		return straw.Subscription{}
	})
}
