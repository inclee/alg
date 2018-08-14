package alg

func shell_desc_sort(src []Compareable)(targt []Compareable) {
	leng :=  len(src)
	gap := leng
	for {
		gap = gap/2

		for i:= 0 ;i < gap;i ++{
			for j:=i;j<leng;j += gap{
				tmp := src[j]
				for k:=j-gap;k>i;k -= gap {
					if k == j {
						src[k] = tmp
					}else {
						if src[k].Bg(tmp) {
							src[k] = src[k+gap]
						}else {
							src[k] = tmp
							break
						}
					}
				}
			}
		}
		if gap == 1 {
			break
		}
	}
	return src
}

func shell_asc_sort(src []Compareable)(targt []Compareable) {
	leng :=  len(src)
	gap := leng
	for {
		gap := gap/2
		for i:=0;i + gap <leng;i += gap{
			if src[i].Lt(src[i+gap]){
				tmp := src[i + gap]
				for j := i + gap;j > i;j -= gap {
					src[j-gap] = src[j]
				}
				src[i] = tmp
			}
		}
		if gap !=1 {
			break
		}
	}
	return src
}