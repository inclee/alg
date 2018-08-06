package alg

import (
	"testing"
	"time"
	"sync"
)

func TestNewFunel(t *testing.T) {
	wt := sync.WaitGroup{}
	funnel := NewFunel(10,0.001, func(v ...interface{}) {
		t.Log("fun time: ",time.Now(),v)
	})
	funnel.Flow(&wt)
	for  i := 0; i< 100000;i++{
		funnel.Append(i)
	}
	wt.Done()
}

