package matrix

type Matrix interface {
	Dims() (r, c int)
	At(i, j int) float64
	Set(i, j int, v float64)
}
