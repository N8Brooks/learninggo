package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	after := time.After(2 * time.Second)
	var idx, sum int
	var reason string
loop:
	for ; ; idx++ {
		select {
		case <-after:
			reason = "timeout"
			break loop
		default:
		}
		num := rand.Intn(100_000_000)
		sum += num
		if num == 1234 {
			reason = "found 1234"
			break
		}
	}
	fmt.Printf("i = %d, sum = %d\n, %s", idx, sum, reason)
}
