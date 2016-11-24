package utils

func Sum(x int) int{
	sum := 0
	for i:=0; i<x; i++ {
		sum = sum + i
	}
	return sum
}

func SquareRoot(num float64) float64 {
	root := num/2
	for i := 0; i < 10 ; i++ {
		root = newtonMethod(num, root)
	}
	return root
}

func newtonMethod(num float64, zn float64) float64 {
	nextVal := zn - (zn*zn - num)/(2*zn)
	return nextVal
}
