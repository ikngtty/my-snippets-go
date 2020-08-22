package main

const N = 1_000_000_007
const MAX = 2_000_001

var fac, finv, inv []int

func main() {
	fac = make([]int, MAX)
	fac[0] = 1
	fac[1] = 1
	finv = make([]int, MAX)
	finv[0] = 1
	finv[1] = 1
	inv = make([]int, MAX)
	inv[1] = 1
	for i := 2; i < MAX; i++ {
		fac[i] = fac[i-1] * i % N
		inv[i] = N - inv[N%i]*(N/i)%N
		finv[i] = finv[i-1] * inv[i] % N
	}
}

func Combi(n, k int) int {
	if n < k {
		return 0
	}
	if n < 0 || k < 0 {
		return 0
	}
	return fac[n] * (finv[k] * finv[n-k] % N) % N
}
