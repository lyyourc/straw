package straw

type (
	NextFunc[T any] func(T)
	ErrorFunc       func(error)
	CompleteFunc    func()
)

type Consumer[T any] interface {
	Next(value T)
	Error(error)
	Complete()
}

// type UnSubcribable interface {
// 	UnSubscribe()
// }

// type Subscribable[T any] interface {
// 	Subscribe(Listenerable[T]) UnSubcribable
// }
type OperatorFunc[T any, R any] func(Observable[T]) Observable[R]
