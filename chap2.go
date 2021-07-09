package main

import (
	"fmt"

	"github.com/tychy/jurljs/tempconv"
)

func main() {
	f := tempconv.Fahrenheit(10.0)
	c := tempconv.Celsius(100.0)
	fmt.Println(tempconv.FToC(f))
	fmt.Println(tempconv.CToK(c))
}
