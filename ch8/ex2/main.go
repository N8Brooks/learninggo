package main

import "fmt"

type Printable interface {
	~int | ~float64
	fmt.Stringer
}

type Int int

func (i Int) String() string {
	return fmt.Sprint("Int: ", int(i))
}

type Float float64

func (f Float) String() string {
	return fmt.Sprint("Float: ", float64(f))
}

func Print[T Printable](x T) {
	fmt.Println(x)
}

func main() {
	Print(Int(42))
	Print(Float(42.0))
}
