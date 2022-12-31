package observable

import (
	"fmt"
	"testing"

	"github.com/lyyourc/straw"
)

func TestOf(t *testing.T) {
	Of(1, 2, 3).Subscribe(straw.Observer[int]{
		Next: func(val int) {
			fmt.Printf("next: %v, ", val)
		},
	})
}
