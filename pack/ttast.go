package pack

import "fmt"

const (
	tta = 7 //需要N瓶
	ttk = 2 //2瓶可以换一瓶
)

//func main() {
//	var blankCount, sum int
//	i := 0
//	blankCount = 0
//	for i <= tta {
//		i++
//		blankCount++
//		sum ++
//		if blankCount%ttk == 0 {
//			i ++ //换一个
//			blankCount = 0
//		}
//	}
//
//	fmt.Println("count:", sum)
//}

func main3() {
	blackCount := 0
	fmt.Println(f1(tta, ttk, &blackCount, true))
}

//主函数
func f1(n int, m int, blackCount *int, first bool) int {
	if n == 1 {
		return 1
	}
	*blackCount++
	if !first && *blackCount%m == 0 {
		*blackCount = 0
		return f1(n-1, m, blackCount, false)
	} else {
		return f1(n-1, m, blackCount, false) + 1
	}
}
