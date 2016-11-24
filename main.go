package main

import (
	"fmt"
	"math"
	"github.com/lsudarshan/hello/utils"
)

func main(){
	//fmt.Println(utils.Hello())
	//fmt.Println(utils.Swap("hello", "world"))
	//fmt.Println("Converting int to string...")
	//fmt.Println(string(12345))
	//fmt.Println(utils.Sum(10))
	num := float64(30)
	fmt.Printf("square root of : %g \n", num)
	fmt.Println(math.Sqrt(num))
	fmt.Println(utils.SquareRoot(num))
}
