package utils

import "fmt"

type Queue []*Node

func (q *Queue) Push(v *Node) {
	*q = append(*q, v)

}

func (q *Queue) Pop() (*Node, error) {
	if len(*q) == 0 {
		return nil, fmt.Errorf("queue is empty")
	}
	v := (*q)[0]

	*q = (*q)[1:]

	return v, nil
}
