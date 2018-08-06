package alg

import "testing"

func TestDequeue(t *testing.T) {
	queue := NewQueue(10)
	err := queue.EnQueue(1,2,3,4,5,6)
	if err != nil {
		t.Fail()
	}
	err = queue.EnQueue(7)
	if err != nil {
		t.Fail()
	}
	err = queue.EnQueue(8,9,10,11)
	if err != nil {
		if err != ErrQueueFull {
			t.Fail()
		}else {
			t.Log(err)
		}
	}
}