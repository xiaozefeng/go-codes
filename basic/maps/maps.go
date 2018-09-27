package main

import "fmt"

func main() {
	m := map[string]string{
		"name":  "jack",
		"age":   "18",
		"hobby": "eat,sleep,code",
	}
	fmt.Println(m)
	m2 := make(map[string]int) // is empty
	fmt.Println(m2)
	var m3 map[string]int // is nil
	fmt.Println(m3)

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("Getting value")
	if name, ok := m["name"]; ok {
		fmt.Println("get value:", name)
	} else {
		fmt.Println("获取不到name")
	}
	fmt.Println("Deleting value ")
	delete(m, "name")
	fmt.Println(m2)
	if name, ok := m["name"]; ok {
		fmt.Println("get value:", name)
	} else {
		fmt.Println("获取不到name")
	}

}
