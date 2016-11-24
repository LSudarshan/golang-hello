package utils

func Hello() string{
	return "hello world!"
}

func Swap(x string, y string) (string,string) {
	a, b := y, x
	return a,b
}