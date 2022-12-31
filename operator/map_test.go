package operator

import (
	"fmt"
	"testing"

	"github.com/lyyourc/straw"
)

func TestMap(t *testing.T) {
	observable := straw.NewObservable(func(sub straw.Subscriber[int]) straw.Subscription {
		sub.Next(42)
		sub.Complete()
		return straw.Subscription{}
	}).Pipe(
		Map(func(value, index int) int {
			return 42 + 1
		}),
	)

	observable.Subscribe(straw.Observer[int]{
		Next: func(i int) {
			fmt.Printf("next: %+v", i)
		},
	})
}
