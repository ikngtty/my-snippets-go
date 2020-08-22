package main

func main() {

}

func Ratio(x, y int) (int, int) {
	if x == 0 {
		if y == 0 {
			return 0, 0
		}
		return 0, 1
	} else if y == 0 {
		return 1, 0
	}

	sign := 1
	if x > 0 && y < 0 {
		sign = -1
	} else if x < 0 && y > 0 {
		sign = -1
	}

	x = Abs(x)
	y = Abs(y)
	var small, big int
	if x < y {
		small, big = x, y
	} else {
		small, big = y, x
	}
	gcd := GCD(small, big)
	return x / gcd, sign * y / gcd
}

func ReverseRatio(x, y int) (int, int) {
	sign := -1
	if x == 0 || y <= 0 {
		sign = 1
	}

	return Abs(y), sign * Abs(x)
}

func GCD(smaller, bigger int) int {
	rem := bigger % smaller
	if rem == 0 {
		return smaller
	}
	return GCD(rem, smaller)
}

func Abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
