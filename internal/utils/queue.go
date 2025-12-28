package utils

import "fmt"

type Queue []int

func (q *Queue) Push(v int) {
	*q = append(*q, v)

}

func (q *Queue) Pop() (int, error) {
	if len(*q) == 0 {
		return 0, fmt.Errorf("queue is empty")
	}
	v := (*q)[0]

	*q = (*q)[1:]

	return v, nil
}
