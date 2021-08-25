package sha256_util

import (
	"crypto/sha256"
	"fmt"

	"github.com/tychy/jurljs/chapter2/popcount"
)

func Diff(s1 string, s2 string) int {
	c1 := sha256.Sum256([]byte(s1))
	c2 := sha256.Sum256([]byte(s2))
	fmt.Printf("%x\n%x\n", c1, c2)
	res := 0
	for i := 0; i < 32; i++ {
		res += popcount.PopCount(uint64(c1[i] ^ c2[i]))
	}
	return res
}
