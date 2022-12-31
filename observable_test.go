package straw

import (
	"fmt"
	"testing"
)

func TestNewObservable(t *testing.T) {
	observable := NewObservable(func(sub Subscriber[int]) Subscription {
		sub.Next(42)
		sub.Complete()
		return Subscription{}
	})

	observable.Subscribe(Observer[int]{
		Next: func(i int) {
			fmt.Printf("next: %+v", i)
		},
	})
}
