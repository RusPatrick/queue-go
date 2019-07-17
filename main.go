package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	errEmptyQueue = errors.New("очередь пустая")
)

type Queue []interface{}

func NewQueue(len int) Queue {
	return make(Queue, len)
}

func (q *Queue) Push(v interface{}) *Queue {
	*q = append(*q, v)
	return q
}

func (q *Queue) Pop() (interface{}, error) {
	if len(*q) != 0 {
		v := (*q)[0]
		*q = (*q)[1:]
		return v, nil
	}
	return nil, errEmptyQueue
}

func (q *Queue) Len() int {
	return len(*q)
}

type Id struct {
	Id int `json:"id"`
}

func main() {
	q := Queue{}
	q.Push("123").Push(123).Push(Id{Id: 112})
	qJson, _ := json.MarshalIndent(q, "", "\t")
	fmt.Println(string(qJson))
	a, err := q.Pop()
	fmt.Println(err)
	qJson, _ = json.MarshalIndent(q, "", "\t")
	fmt.Println(string(qJson))
	fmt.Println(a)
}
