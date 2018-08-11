package alg

type Compareable interface {
	bg(c Compareable)bool
	lt(c Compareable)bool
	eq(c Compareable)bool
}
