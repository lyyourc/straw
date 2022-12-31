package observable

import (
	"fmt"
	"testing"

	"github.com/lyyourc/straw"
	operator "github.com/lyyourc/straw/operator"
)

func TestPipe(t *testing.T) {
	Pipe(
		Of(42),
		operator.Map[int, int](func(value, index int) int {
			return value + 1
		}),
	).
		Subscribe(straw.Observer[int]{
			Next: func(val int) {
				fmt.Printf("next: %v, ", val)
			},
		})
}

func TestPipe2(t *testing.T) {
	Pipe(
		Of(42),
		operator.SwitchMap(func(value, index int) straw.Observable[int] {
			return Of(value + 1)
		}),
	).
		Subscribe(straw.Observer[int]{
			Next: func(val int) {
				fmt.Printf("next: %v, ", val)
			},
		})
}
