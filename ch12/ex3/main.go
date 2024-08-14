package main

import (
	"fmt"
	"math"
	"sync"
)

var map1 func() map[int]float64 = sync.OnceValue(func() map[int]float64 {
	n := 100_0000
	map1 := make(map[int]float64, n)
	for i := 0; i < n; i++ {
		map1[i] = math.Sqrt(float64(i))
	}
	return map1
})

func main() {
	for i := 0; i < 100_000; i += 1_000 {
		fmt.Println(map1()[i])
	}
}
