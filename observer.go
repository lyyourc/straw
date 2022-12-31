package straw

// implement Listenable interface
// friendly to call subscribe function
type Observer[T any] struct {
	Next     func(T)
	Error    func(error)
	Complete func()
}
