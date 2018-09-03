package alg

import (
	"time"
	"sync"
)


type Funnel struct {
	capcity uint
	leftOpt uint
	rat float64
	lastLeedTm time.Time
	queue QueueI
	fundel func(v ...interface{})
	ticker *time.Ticker
	puse chan int
}

func NewFunel(capcity uint,rat float64,funnel func(v ...interface{}))(*Funnel){
	return &Funnel{
		capcity:capcity,
		leftOpt:capcity,
		rat:rat,
		queue:NewQueue(capcity*10),
		fundel:funnel,
		ticker:time.NewTicker( time.Duration(1000*rat)*time.Second),
		puse:make(chan int,1),
	}
}

func (f *Funnel)Flow(group *sync.WaitGroup)(){
	group.Add(1)
	go func() {
		defer group.Done()
		//f.puse <- 1
		for {
			//<- f.puse
			<- f.ticker.C
			v,err := f.queue.DeQueue()
			if err != nil {
				if err == ErrQueueEmpty {
					// 堵塞
				}else {

				}
			}else {
				f.leftOpt += 1
				f.lastLeedTm = time.Now()
				f.fundel(v)
				if f.queue.Count() > 0{
					//f.puse <- 1
				}
			}
		}
	}()
}

func (f *Funnel)Append(v ...interface{})  {
	if f.leftOpt > 0 {
		taked := MinInt(int(f.leftOpt),len(v))
		if f.fundel != nil{
			f.leftOpt -= uint(taked)
			f.fundel(v[0:taked]...)
		}
		if taked != len(v) {
			f.queue.EnQueue(v[taked:]...)
			//f.puse <- 1
		}
	}else {
		f.queue.EnQueue(v...)
	}
}

