package main

import "fmt"

const (
	Low = iota
	Medium
	High
)

type PriorityQueue[PriorTyp comparable, ValTyp any] struct {
	items      map[PriorTyp][]ValTyp
	priorities []PriorTyp
}

func (pq *PriorityQueue[PriorTyp, ValTyp]) Add(priority PriorTyp, value ValTyp) {
	pq.items[priority] = append(pq.items[priority], value)
}

func (pq *PriorityQueue[PriorTyp, ValTyp]) Next() (ValTyp, bool) {
	for i := 0; i < len(pq.priorities); i++ {
		priority := pq.priorities[i]
		items := pq.items[priority]
		if len(items) > 0 {
			next := items[0]
			pq.items[priority] = items[1:]
			return next, true
		}
	}
	var empty ValTyp
	return empty, false
}

func NewPriorityQueue[PriorTyp comparable, ValTyp any](priorities []PriorTyp) PriorityQueue[PriorTyp, ValTyp] {
	return PriorityQueue[PriorTyp, ValTyp]{items: make(map[PriorTyp][]ValTyp), priorities: priorities}
}

func main() {
	queue := NewPriorityQueue[int, string]([]int{High, Medium, Low})
	queue.Add(Low, "L-1")
	queue.Add(High, "H-1")
	fmt.Println(queue.Next())

	queue.Add(Medium, "M-1")
	queue.Add(High, "H-2")
	queue.Add(High, "H-3")

	fmt.Println(queue.Next())
	fmt.Println(queue.Next())
	fmt.Println(queue.Next())
	fmt.Println(queue.Next())
	fmt.Println(queue.Next())
	fmt.Println(queue.Next())
	fmt.Println(queue.Next())
}
