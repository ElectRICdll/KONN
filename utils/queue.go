package utils

import (
	"container/list"
)

type Queue struct {
	queue     *list.List
	maxLength int
}

func NewQueue(max int) Queue {
	return Queue{
		queue:     list.New(),
		maxLength: max,
	}
}

func (q *Queue) Push(obj interface{}) {
	if q.queue.Len() >= q.maxLength {
		Logger.Error("No enough space for new member.")
	} else {
		q.queue.PushBack(obj)
	}
}

func (q *Queue) Poll() interface{} {
	defer q.queue.Remove(q.queue.Front())
	return q.queue.Front().Value
}

func (q *Queue) Remove(index int) {
	for iter := q.queue.Front(); iter != nil; iter.Next() {
		if index == 0 {
			q.queue.Remove(iter)
		}
		index--
	}
}

func (q *Queue) MaxLength() int {
	return q.maxLength
}

func (q *Queue) SetMaxLength(length int) {
	q.maxLength = length
}
