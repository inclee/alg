package alg

import (
	"testing"
	"fmt"
)

func TestShellSort(t *testing.T){
	sortItem := []Compareable{SortItem(27),SortItem(31),SortItem(4),SortItem(123),SortItem(23),SortItem(32),SortItem(365),SortItem(38),SortItem(87)}
	sortedItem := shell_desc_sort(sortItem)
	for _,item := range sortedItem  {
		fmt.Println(item)
	}
}
