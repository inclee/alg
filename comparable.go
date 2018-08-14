package alg

type Compareable interface {
	Bg(c Compareable)bool
	Lt(c Compareable)bool
	Eq(c Compareable)bool
}

type SortItem int

func (s SortItem)Lt(other Compareable)bool {
	return int(s) < int(other.(SortItem))
}
func (s SortItem)Bg(other Compareable)bool {
	return int(s) > int(other.(SortItem))
}

func (s SortItem)Eq(other Compareable)bool {
	return int(s) == int(other.(SortItem))
}