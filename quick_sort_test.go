package alg

import (
	"testing"
	"fmt"
)

func TestQuickSortTest(t *testing.T)  {
	sortItem := []Compareable{SortItem(22),SortItem(13),SortItem(4),SortItem(123),SortItem(23),SortItem(32),SortItem(365),SortItem(38),SortItem(87)}
	sortedItem := quick_desc_sort(sortItem)
	for _,item := range sortedItem  {
		fmt.Println(item)
	}
}
