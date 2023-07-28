package utils

import (
	"container/list"
)

type Queue struct {
	queue     list.List
	maxLength int
}

func NewQueue() {

}

func (q *Queue) Push(obj interface{}) {
	if q.queue.Len() >= q.maxLength {

	}
}

func (q *Queue) Poll() {

}

func (q *Queue) toList() list.List {
	return list.List{}
}

func (q *Queue) MaxLength() {

}

func (q *Queue) SetMaxLength() {

}
