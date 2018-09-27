package main

import "fmt"

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{2, 3, 4, 5, 6, 7}
	fmt.Printf("arr1:%v, arr2:%+v, arr3:%#v\n", arr1, arr2, arr3)
	var grids [4][5]int
	fmt.Printf("%v\n", grids)
	fmt.Printf("grids.length: %d\n", len(grids))
	for i := 0; i < len(grids); i++ {
		for j := 0; j < len(grids[i]); j++ {
			grids[i][j] = j
		}
	}
	fmt.Printf("%v\n", grids)
	fmt.Println("range:")
	for i, val := range arr3 {
		fmt.Printf("%d,%d\n", i, val)
	}

	fmt.Println("invoke printArray")
	printArray(arr3)
	fmt.Println(arr3)
}

func printArray(arr [6]int) {
	arr[0] = 100
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}
