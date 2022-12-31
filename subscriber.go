package straw

type Subscriber[T any] struct {
	onNext     NextFunc[T]
	onError    ErrorFunc
	onComplete CompleteFunc
}

type SubscriberOptions[T any] func(*Subscriber[T])

func NewSubscriber[T any](onNext func(T), onError func(error), onComplete func()) Subscriber[T] {
	return Subscriber[T]{
		onNext:     onNext,
		onError:    onError,
		onComplete: onComplete,
	}
}

func (subscriber Subscriber[T]) Next(value T) {
	if subscriber.onNext != nil {
		subscriber.onNext(value)
	}
}

func (subscriber Subscriber[T]) Error(err error) {
	if subscriber.onError != nil {
		subscriber.onError(err)
	}
}

func (subscriber Subscriber[T]) Complete() {
	if subscriber.onComplete != nil {
		subscriber.onComplete()
	}
}
