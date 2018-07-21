package array

func Swap(num []int, i, j int) {
	num[i] = num[i] ^ num[j]
	num[j] = num[i] ^ num[j]
	num[i] = num[i] ^ num[j]
}
