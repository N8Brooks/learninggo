package main

import (
	"fmt"
)

func UpdateSlice(arr []string, s string) {
	arr[len(arr)-1] = s
	fmt.Println(arr)
}

func GrowSlice(arr []string, s string) {
	arr = append(arr, s)
	fmt.Println(arr)

}

func main() {
	s1 := []string{"a", "b", "c", "d"}
	UpdateSlice(s1, "z")
	fmt.Println(s1)

	s2 := []string{"a", "b", "c", "d"}
	GrowSlice(s2, "z")
	fmt.Println(s2)
}
