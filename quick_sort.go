package alg

func quick_desc_sort(src []Compareable)(targt []Compareable) {
	if len(src) > 1{
		base_idx := 0
		base := src[0]
		for i:= 1 ;i < len(src);i++ {
			if base.Lt(src[i]){
				tmp := src[i]
				for j:= i;j>base_idx;j--{
					src[j] = src[j-1]
				}
			    src[base_idx] = tmp
			    base_idx +=1
			}
		}
		pre := src[0:base_idx]
		bak := src[base_idx+1:len(src)]
		quick_desc_sort(pre)
		quick_desc_sort(bak)
	}
	return src
}


func quick_asc_sort(src []Compareable)(targt []Compareable) {
	if len(src) > 1{
		base_idx := 0
		base := src[0]
		for i:= 1 ;i < len(src);i++  {
			if base.Bg(src[i]){
				tmp := src[i]
				for j:= i;j>base_idx;j--{
					src[j] = src[j-1]
				}
				src[base_idx] = tmp
				base_idx +=1
			}
		}
		pre := src[0:base_idx]
		bak := src[base_idx+1:len(src)]
		quick_asc_sort(pre)
		quick_asc_sort(bak)
	}
	return src
}