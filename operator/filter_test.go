package operator

import (
	"fmt"
	"testing"

	"github.com/lyyourc/straw"
)

func TestFilter(t *testing.T) {
	obs := straw.NewObservable(func(sub straw.Subscriber[int]) straw.Subscription {
		for _, item := range []int{42, 10} {
			sub.Next(item)
		}
		sub.Complete()
		return straw.Subscription{}
	}).Pipe(
		Filter(func(val, index int) bool {
			return val > 1
		}),
	)

	obs.Subscribe(straw.Observer[int]{
		Next: func(i int) {
			fmt.Printf("next: %+v \n", i)
		},
	})
}
