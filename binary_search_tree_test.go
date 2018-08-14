package alg

import (
	"testing"
	"fmt"
)

func TestBSTree(t *testing.T)  {
	sortItem := []Compareable{SortItem(2),SortItem(3),SortItem(4),SortItem(123),SortItem(23),SortItem(32),SortItem(365),SortItem(38),SortItem(87)}
	bst := new(BSTree)
	for _,i := range sortItem{
		node := new(Node)
		node.value = i
		bst.Insert(node)
	}
	sortedItem := bst.InOrderTraversal()
	for _,item := range sortedItem  {
		fmt.Println(item)
	}
	fined := bst.Find(SortItem((4)))
	if fined == nil {
		t.Log("can't find")
	}else {
		t.Logf("find object %d",fined.value)
	}
}
