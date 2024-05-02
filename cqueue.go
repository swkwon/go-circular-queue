package cqueue

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
)

var ErrQueueIsFull = errors.New("queue is full")
var ErrQueueIsEmpty = errors.New("queue is empty")

type CQueue[T any] struct {
	container []T
	size      int
	front     int
	rear      int
	cnt       int
	mutex     *sync.Mutex
}

// MakeQueue is a generic function that creates a circular queue.
// The parameter sets the size of the queue.
func MakeQueue[T any](size int) *CQueue[T] {
	if size <= 0 {
		size = 100
	}
	return &CQueue[T]{
		container: make([]T, size),
		size:      size,
		mutex:     &sync.Mutex{},
	}
}

// IsFull method checks to see if the queue is full.
func (g *CQueue[T]) IsFull() bool {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	return g.isFullImpl()
}

func (g *CQueue[T]) isFullImpl() bool {
	return g.cnt >= g.size
}

// IsEmpty method checks to see if the queue is empty.
func (g *CQueue[T]) IsEmpty() bool {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	return g.isEmptyImpl()
}

func (g *CQueue[T]) isEmptyImpl() bool {
	return g.cnt <= 0
}

// Push method inserts an element into the queue.
func (g *CQueue[T]) Push(v T) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	if g.isFullImpl() {
		return ErrQueueIsFull
	}
	g.container[g.rear] = v
	g.rear = (g.rear + 1) % g.size
	g.cnt++
	return nil
}

// Pop method fetches the first element in the queue.
func (g *CQueue[T]) Pop() (T, error) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	var ret T
	if g.isEmptyImpl() {
		return ret, ErrQueueIsEmpty
	}
	ret = g.container[g.front]
	g.front = (g.front + 1) % g.size
	g.cnt--
	return ret, nil
}

// View method returns all the elements in the queue as an array of strings in json format.
// You can create String() methods for the elements to output them in the formatting you want.
func (g *CQueue[T]) View() string {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	var strList []string
	if g.isEmptyImpl() {
		return "[]"
	}

	i := g.front

	for {
		strList = append(strList, fmt.Sprint(g.container[i]))
		i = (i + 1) % g.size
		if i == g.rear {
			break
		}
	}

	b, e := json.Marshal(&strList)
	if e != nil {
		return "[]"
	}
	return string(b)
}
