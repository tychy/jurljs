package main

import (
	"fmt"

	"github.com/tychy/jurljs/chapter2/popcount"
	"github.com/tychy/jurljs/chapter2/tempconv"
)

func temp() {
	f := tempconv.Fahrenheit(10.0)
	c := tempconv.Celsius(100.0)
	fmt.Println(tempconv.FToC(f))
	fmt.Println(tempconv.CToK(c))
}
func pc() {
	fmt.Println(popcount.PopCount(100))     //64 + 32 + 4
	fmt.Println(popcount.PopCountLoop(100)) //64 + 32 + 4

	fmt.Println(popcount.PopCount(101))     //64 + 32 + 4
	fmt.Println(popcount.PopCountLoop(101)) //64 + 32 + 4
}

func main() {
	temp()
	pc()
}
