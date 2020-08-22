package main

import (
	"fmt"
)

func main() {
	fmt.Println(Factorize(792))
	// -> [1 2 4 8 3 6 12 24 9 18 36 72 11 22 44 88 33 66 132 264 99 198 396 792]
}

func Factorize(n int) []int {
	primeCounts := PrimeFactorize(n)
	primeLen := len(primeCounts)
	if primeLen == 0 {
		return []int{1}
	}

	factorsLen := 1
	for _, primeCount := range primeCounts {
		factorsLen *= primeCount.Count + 1
	}

	factors := make([]int, 0, factorsLen)
	usePrimeCounts := make([]int, primeLen)
	lastPrimeCount := primeCounts[primeLen-1].Count
	for usePrimeCounts[primeLen-1] <= lastPrimeCount {
		factor := 1
		for i, primeCount := range primeCounts {
			factor *= Pow(primeCount.Value, usePrimeCounts[i])
		}
		factors = append(factors, factor)

		usePrimeCounts[0]++
		// carry
		for i := 0; i < primeLen-1; i++ {
			if usePrimeCounts[i] > primeCounts[i].Count {
				usePrimeCounts[i] = 0
				usePrimeCounts[i+1]++
			} else {
				break
			}
		}
	}

	return factors
}

func PrimeFactorize(n int) []IntTally {
	primeCounts := make([]IntTally, 0)

	addCountIfExist := func(i int) {
		count := 0
		for n%i == 0 {
			count++
			n /= i
		}
		if count > 0 {
			primeCounts = append(primeCounts, IntTally{i, count})
		}
	}

	addCountIfExist(2)
	addCountIfExist(3)

	flag := true
	for i := 5; n > 1 && i*i <= n; {
		addCountIfExist(i)

		if flag {
			i += 2
		} else {
			i += 4
		}
		flag = !flag
	}

	if n > 1 {
		primeCounts = append(primeCounts, IntTally{n, 1})
	}

	return primeCounts
}

type IntTally struct {
	Value int
	Count int
}

func Pow(base int, exponent int) int {
	if exponent < 0 {
		panic(fmt.Sprintf("exponent: %d", exponent))
	}
	answer := 1
	for i := 0; i < exponent; i++ {
		answer *= base
	}
	return answer
}
