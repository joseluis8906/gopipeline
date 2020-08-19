package completion

import (
	"fmt"
	"gopipeline/order"
	"sync"
	"time"

	"syreclabs.com/go/faker"
)

func Complete(ods <-chan order.Order) <-chan order.Order {
	out := make(chan order.Order)
	go func() {
		var wg sync.WaitGroup
		for od := range ods {
			wg.Add(1)
			go func(ord order.Order) {
				out <- fetchData(ord)
				wg.Done()
			}(od)
		}

		wg.Wait()
		close(out)
	}()

	return out
}

func fetchData(od order.Order) order.Order {
	parts := []string{
		"elevation",
		"toolkit",
		"position",
	}
	in := make(chan string, len(parts))
	for _, part := range parts {
		go func(in chan<- string, part string) {
			time.Sleep(150 * time.Millisecond)
			if faker.RandomInt(1, 100) > 95 {
				in <- "error"
			} else {
				in <- part
			}
		}(in, part)
	}

	for i := 0; i < len(parts); i++ {
		resp := <-in
		od.State = fmt.Sprintf("%s, %s", od.State, resp)
	}

	close(in)
	return od
}
