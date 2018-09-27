package main

import (
	"fmt"
)

func main() {
	var s []int
	for i := 0; i < 100; i++ {
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8, 10}
	printSlice(s1, "s1")

	s2 := make([]int, 16)
	printSlice(s2, "s2")

	s3 := make([]int, 10, 30)
	printSlice(s3, "s3")
	fmt.Println("Coping slice")
	copy(s2, s1)
	printSlice(s2, "s2")
	fmt.Println("Deleting element from slice")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2, "s2")
	fmt.Println("Poping from front")
	front := s2[0]
	s2 = s2[1:]
	printSlice(s2, "s2")
	fmt.Println("front", front)
	fmt.Println("Poping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	printSlice(s2, "s2")
	fmt.Println("tail", tail)
	printSlice(s2, "s2")

}

func printSlice(s []int, name string) {
	fmt.Printf("slice: %s, len: %d, cap: %d, value: %v\n", name, len(s), cap(s), s)
}
