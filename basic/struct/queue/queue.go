package queue

// slice 别名
type Queue []int

func (q *Queue) Push(e int) {
	*q = append(*q, e)
}

func (q *Queue) Pop() int {
	ret := (*q)[0]
	*q = (*q)[1:]
	return ret
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0

}
