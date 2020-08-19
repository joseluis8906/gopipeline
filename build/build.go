package build

import (
	"gopipeline/completion"
	"gopipeline/fetching"
	"gopipeline/order"
	"gopipeline/validation"
)

func Build() <-chan order.Order {
	fetchedOrders := fetching.Fetch(500000)
	completedOrders := completion.Complete(fetchedOrders)
	return validation.Validate(completedOrders)
}
