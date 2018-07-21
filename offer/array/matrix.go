package array

import "fmt"

//题目：输入一个矩阵，按照从外向里以顺时针的顺序依次扫印出每一个数字
//解题思路：
//把打印一圈分为四步：第一步从左到右打印一行，第二步从上到下打印一列，第三步从右到左打印一行，第四步从下到上打印一列。
// 每一步我们根据起始坐标和终止坐标用一个循环就能打印出一行或者一列。
// 不过值得注意的是，最后一圈有可能退化成只有一行、只有一列，甚至只有一个数字，因此打印这样的一圈就不再需要四步。
// 图4.4 是几个退化的例子， 打印一圈分别只需要三步、两步甚至只有一步。
//因此我们要仔细分析打印时每一步的前提条件。第一步总是需要的， 因为打印一圈至少有一步。如果只有一行，那么就不用第二步了。也就是需要
//第二步的前提条件是终止行号大于起始行号。需要第三步打印的前提条件是圈内至少有两行两列，也就是说除了要求终止行号大于起始行号之外，还要求终止列号大于起始列号。
//同理，需要打印第四步的前提条件是至少有三行两列，因此要求终止行号比起始行号至少大2 ， 同时终止列号大于起始列号。
func PrintMatrixByCircle(a [][]int, handler func(v int)) {
	m := len(a) - 1
	n := len(a[0]) - 1
	loopX, loopY := m/2, n/2 //x,y为循环次数
	for x, y := 0, 0; x <= loopX && y <= loopY; {
		fmt.Println("\n----", x, y)
		printMatrixByCircleInner(a, x, y, handler)
		x++
		y++
	}
}

func printMatrixByCircleInner(a [][]int, x, y int, handler func(v int)) {
	n := len(a) - 1
	m := len(a[0]) - 1
	fmt.Println("1圈")
	//1  a[x][y] -->a[x][m-y] j
	for j := y; j <= m-y; j++ {
		handler(a[x][j])
	}
	fmt.Println("\n2圈")
	// 2 a[x+1][m-y]-->a[n-x][m-y]
	if x+1 <= n-x && m-y > y {
		for i := x + 1; i <= n-x; i++ {
			handler(a[i][m-y])
		}
	}
	fmt.Println("\n3圈")
	// 3 a[n-x][m-y-1] -->a[n-x][y]
	if m-y-1 >= y && n-x > x {
		for j := m - y - 1; j >= y; j-- {
			handler(a[n-x][j])
		}
	}
	fmt.Println("\n4圈")
	//4 a[n-x-1][y] --> a[x+1][y]
	if n-x-1 >= x+1 && m-y > y {
		for i := n - x - 1; i >= x+1; i-- {
			handler(a[i][y])
		}
	}

}
