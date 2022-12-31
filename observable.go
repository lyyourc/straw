package straw

type Observable[T any] struct {
	produce func(ob Subscriber[T]) Subscription
}

func NewObservable[T any](produce func(Subscriber[T]) Subscription) Observable[T] {
	return Observable[T]{
		produce: produce,
	}
}

func (o Observable[T]) Subscribe(ob Observer[T]) Subscription {
	subscriber := NewSubscriber(ob.Next, ob.Error, ob.Complete)
	subscription := o.produce(subscriber)
	return subscription
}

func (o Observable[any]) Pipe(fns ...OperatorFunc[any, any]) Observable[any] {
	current := o

	for _, fn := range fns {
		current = fn(current)
	}

	return current
}
