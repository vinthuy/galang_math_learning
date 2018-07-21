package number

func GetOneCount(a int) int {
	count := 0
	n := a
	for n != 0 {
		n = (n - 1) & n
		count ++
	}
	return count
}
func Power(base int, exponent int64) float64 {
	if base == 0 && exponent == 0 {
		panic("invalid input!!")
	}

	if exponent == 0 {
		return 1
	}
	exp := exponent

	if exp < 0 {
		exp = -exp
	}
	result := powerWithUnsignedExponent(base, exp)
	// 指数是负数，要进行求倒数
	if exponent < 0 {
		result = 1 / result
	}
	return result
}
func powerWithUnsignedExponent(base int, exponent int64) float64 {
	// 如果指数为0，返回1
	if exponent == 0 {
		return 1
	}

	// 指数为1，返回底数
	if exponent == 1 {
		return float64(base)
	}

	// 如果是偶数 a[n] = a[n/2]*a[n/2]
	// 如果是奇数 a[n] = a[(n-1)/2] * a[(n-1)/2] *a
	//  其中[]代表指数
	result := powerWithUnsignedExponent(base, exponent>>1)
	result = result * result
	if exponent%2 != 0 {
		result = result * float64(base)
	}
	return result
}
