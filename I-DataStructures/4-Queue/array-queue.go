package main

import "fmt"

type ArrayQueue struct {
	data     []string
	front    int
	rear     int
	capacity int
}

func createQueue(capacity int) ArrayQueue {
	queue := ArrayQueue{}
	queue.data = make([]string, capacity)
	queue.front = 0
	queue.rear = -1
	return queue
}

func (q *ArrayQueue) isEmpty() bool {
	return q.front > q.rear
}

func (q *ArrayQueue) Enqueue(v string) {
	q.rear++
	q.data[q.rear] = v
}

func (q *ArrayQueue) Dequeue() string {
	if !q.isEmpty() {
		v := q.data[q.front]
		q.front++
		return v
	}
	return ""
}

func (q *ArrayQueue) Display() {
	for i := q.front; i <= q.rear; i++ {
		fmt.Println(q.data[i])
	}
}

func main() {
	queue := createQueue(5)
	queue.Enqueue("Hello")
	queue.Enqueue("Hello2")
	queue.Enqueue("Hello3")
	queue.Dequeue()
	queue.Display()
}
