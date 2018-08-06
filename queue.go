package alg

import (
	"sync"
	"github.com/pkg/errors"
)

type QueueI interface {
	EnQueue(value ...interface{})(error)
	DeQueue()(value interface{},err error)
	Count()(int)
}

var  (
	ErrQueueFull = errors.New("queue is full")
	ErrQueueEmpty = errors.New("queue is empty")
)

func NewQueue(capcity uint)(QueueI) {
	return &Queue{
		buf:make([]interface{},0,capcity),
		capcity:capcity,
		mtx:sync.Mutex{},
	}
}

type Queue struct {
	buf []interface{}
	capcity uint
	mtx sync.Mutex
}

func(q *Queue)EnQueue(value ...interface{})(error){
	defer  q.mtx.Unlock()
	q.mtx.Lock()
	if int(q.capcity) <= len(q.buf) + len(value){
		return ErrQueueFull
	}
	q.buf = append(q.buf, value...)
	return nil
}

func(q *Queue)DeQueue()(value interface{},err error){
	defer  q.mtx.Unlock()
	q.mtx.Lock()
	if len(q.buf) == 0 {
		err = ErrQueueEmpty
	}else {
		value = q.buf[0]
		q.buf = q.buf[1:len(q.buf)]
	}
	return value,err
}
func (q *Queue)Count()(int){
	defer  q.mtx.Unlock()
	q.mtx.Lock()
	return len(q.buf)
}

