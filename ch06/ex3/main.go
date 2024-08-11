package main

import (
	"fmt"
	"time"
)

type Person struct {
	firstName string
	lastName  string
	age       int
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

// run with > GODEBUG=gctrace=1 GOGC=100 go run main.go
func main() {
	defer timer("main")()
	n := 10_000_000
	persons := []Person{}
	for i := 0; i < n; i++ {
		persons = append(persons, Person{"John", "Doe", 30})
	}
	println(&persons)
}
