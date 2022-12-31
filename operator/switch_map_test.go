package operator

import (
	"fmt"
	"testing"

	"github.com/lyyourc/straw"
)

func TestSwitchMap(t *testing.T) {
	plusOneObservable := func(value int) straw.Observable[int] {
		return straw.NewObservable[int](func(s straw.Subscriber[int]) straw.Subscription {
			s.Next(value + 1)
			return straw.Subscription{}
		})
	}

	observable := straw.NewObservable(func(sub straw.Subscriber[int]) straw.Subscription {
		sub.Next(42)
		sub.Complete()
		return straw.Subscription{}
	}).Pipe(
		SwitchMap[int, int](func(val, index int) straw.Observable[int] {
			return plusOneObservable(val)
		}),

		Map(func(value, index int) int {
			return value * 2
		}),
	)

	observable.Subscribe(straw.Observer[int]{
		Next: func(i int) {
			fmt.Printf("next: %+v", i)
		},
	})
}
