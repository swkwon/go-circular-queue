# go-circular-queue
go-circular-queue는 thread safety한 generic circular queue 입니다.
# 패키지 설치
```
$ go get github.com/swkwon/go-circular-queue@latest
```
# 시작하기
```
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
	q := cqueue.New[*Job](3)
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
```
```
    cqueue.New[int](3)
```
queue를 만듭니다. `[*Job]`는 큐의 데이터타입입니다. int 형의 데이터가 보관됩니다. `(3)`은 큐의 크기를 나타냅니다. 
```
q.Push(1)
```
`Push` 메서드를 이용하여 큐의 컨테이너에 데이터를 보관합니다. 오버플로우 시 error를 리턴합니다.
```
ret, err = q.Pop()
```
맨앞의 데이터를 꺼내기 위해서 `Pop` 메서드를 호출합니다. 비어있는 큐일 경우 error를 리턴합니다.