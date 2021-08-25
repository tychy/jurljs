package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		// iを右シフトしたらi/2になるので、逆にi/2を左シフトしたもの+iの右端のビットを知れば良い
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
	res := 0
	for i := 0; i < 256; i++ {
		res += int(byte(x>>i) & 1)
	}
	return res
}
