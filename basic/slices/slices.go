package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("arr[2:6]", arr[2:6])
	fmt.Println("arr[:6]", arr[:6])
	s1 := arr[2:]
	fmt.Println("s1:", s1)
	s2 := arr[:]
	fmt.Println("s2:", s2)

	fmt.Println("after updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)
	fmt.Println("after updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)
	fmt.Println("Reslice")
	s2 = s2[:5]
	fmt.Println("s2[:5]", s2)
	s2 = s2[2:]
	fmt.Println("s2[2:]", s2)
}

func updateSlice(sl []int) {
	sl[0] = 100
}
