package main

import (
	"fmt"
	"log"

	cqueue "github.com/swkwon/go-circular-queue"
)

func print(err error) {
	if err != nil {
		log.Println(err)
	}
}

type Job struct {
	Name string
	ID   int
}

func (j *Job) String() string {
	return fmt.Sprintf("name=%s,id=%d", j.Name, j.ID)
}

func main() {
	q := cqueue.MakeQueue[*Job](3)
	var ret *Job
	var err error
	err = q.Push(&Job{Name: "one", ID: 1})
	print(err)
	err = q.Push(&Job{Name: "two", ID: 2})
	print(err)
	err = q.Push(&Job{Name: "three", ID: 3})
	print(err)
	err = q.Push(&Job{Name: "four", ID: 4})
	print(err)
	log.Println(q.View())
	ret, err = q.Pop()
	log.Println("pop: ", ret, err)
	log.Println(q.View())
	ret, err = q.Pop()
	log.Println("pop: ", ret, err)
	log.Println(q.View())
	ret, err = q.Pop()
	log.Println("pop: ", ret, err)
	log.Println(q.View())
	ret, err = q.Pop()
	log.Println("pop: ", ret, err)
	log.Println(q.View())
}
