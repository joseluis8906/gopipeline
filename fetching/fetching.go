package fetching

import (
	"gopipeline/order"
)

func Fetch(max int) <-chan order.Order {
	out := make(chan order.Order)
	go func() {
		for i := 0; i < max; i++ {
			out <- order.Order{ID: i + 1}
		}

		close(out)
	}()

	return out
}
