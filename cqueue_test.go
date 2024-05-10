package cqueue

import (
	"fmt"
	"log"
	"testing"
)

func TestGQueue(t *testing.T) {
	q := New[int](10)
	for i := 0; i <= 10; i++ {
		err := q.Push(i)
		log.Println(fmt.Sprintf("i=%d", i), err)
	}
	for i := 0; i <= 10; i++ {
		ret, err := q.Pop()
		log.Println(ret, err)
	}

}
