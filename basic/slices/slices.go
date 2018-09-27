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
	fmt.Println("Extending slice:")
	s1 = arr[2:6]
	printSlice(s1, "s1")
	s2 = s1[3:5]
	printSlice(s2, "s2")
	s3 := append(s2, 10)
	printSlice(s3, "s3")
	s4 := append(s3, 11)
	// s4 and s5 no longer view arr
	printSlice(s4, "s4")
	s5 := append(s4, 12)
	printSlice(s5, "s5")
}

func printSlice(s []int, name string) {
	fmt.Printf("%s:%v, len:%d, cap:%d\n", name, s, len(s), cap(s))
}

func updateSlice(sl []int) {
	sl[0] = 100
}
