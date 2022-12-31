package test

import (
	"fmt"
	"testing"

	"github.com/lyyourc/straw"
	"github.com/lyyourc/straw/observable"
	"github.com/lyyourc/straw/operator"
)

func TestObservablePipe(t *testing.T) {
	s := observable.Of(42)

	plusOne := func(val int) straw.Observable[int] {
		return straw.NewObservable(func(sub straw.Subscriber[int]) straw.Subscription {
			sub.Next(val + 1)
			sub.Complete()
			return straw.Subscription{}
		})
	}

	s.Pipe(
		operator.SwitchMap(func(val int, index int) straw.Observable[int] {
			return plusOne(val)
		}),
		operator.Map(func(value int, index int) int {
			return value + 1
		}),
	).Subscribe(straw.Observer[int]{
		Next: func(i int) {
			fmt.Printf("next: %+v\n", i)
		},
	})
}
