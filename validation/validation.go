package validation

import (
	"gopipeline/order"
	"strings"
)

func Validate(ods <-chan order.Order) <-chan order.Order {
	out := make(chan order.Order)
	go func() {
		for od := range ods {
			if !strings.Contains(od.State, "error") {
				out <- od
			}
		}

		close(out)
	}()

	return out
}
