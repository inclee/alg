package alg

func MinInt(a int,b int) (int){
	if a > b {
		return b
	}
	return a
}

func Min(a Compareable,b Compareable)(v Compareable){
	if a.Lt(b){
		return a
	}
	return b
}

func Max(a Compareable,b Compareable)(v Compareable){
	if a.Bg(b){
		return a
	}
	return b
}

func Middle(count int)(idx int){
	if count %2 == 0 {
		idx = count/2
	}else {
		idx = count/2 + 1
	}
	return idx
}

