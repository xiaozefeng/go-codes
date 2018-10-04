package adder

// 闭包
func Adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}
