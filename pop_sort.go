package alg

func pop_desc_sort(src []Compareable)(tagt []Compareable){
	leng := len(src)
	for i := 0 ; i < leng;i++{
		for j := i; j<leng; j++{
			if src[i].lt(src[j]) {
				src[i],src[j] = src[j],src[i]
			}
		}
	}
	return src
}

func pop_asc_sort(src []Compareable)(tagt []Compareable){
	leng := len(src)
	for i := 0 ; i < leng;i++{
		for j := i; j<leng; j++{
			if src[i].bg(src[j]) {
				src[i],src[j] = src[j],src[i]
			}
		}
	}
	return src
}
