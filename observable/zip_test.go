package observable

import (
	"fmt"
	"testing"

	"github.com/lyyourc/straw"
)

func TestZip(t *testing.T) {
	observable := straw.NewObservable(func(sub straw.Subscriber[int]) straw.Subscription {
		sub.Next(42)
		sub.Complete()
		return straw.Subscription{}
	})
	observable2 := straw.NewObservable(func(s straw.Subscriber[int]) straw.Subscription {
		s.Next(24)
		return straw.Subscription{}
	})

	Zip(observable, observable2).Subscribe(straw.Observer[[]int]{
		Next: func(i []int) {
			fmt.Printf("next: %+v", i)
		},
	})
}
