package main

import (
	"fmt"
	"gopipeline/build"
	"os"
	"time"
)

func main() {
	start := time.Now()
	counter := 0
	for range build.Build() {
		counter++
	}

	fmt.Println(counter)
	elapsed := time.Since(start)
	fmt.Printf("in %s", elapsed)
	os.Exit(0)
}
