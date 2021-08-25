package main

import (
	"fmt"

	"github.com/tychy/jurljs/chapter4/sha256_util"
)

func main() {
	fmt.Println(sha256_util.Diff("x", "X"))
	fmt.Println(sha256_util.Diff("x", "x"))
}
