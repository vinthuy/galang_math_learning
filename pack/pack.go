package pack

import (
	"fmt"
	"math"
)

const (
	n = 6  // n件物品
	c = 13 // 背包容量
	N = 15
)

var x [n + 1]int                     //表示0，1背包问题
var m [N][N]int                      //m[i][j]表示面对第i件物品 ，背包容量为j的的最大价值数
var v = [N]int{0, 8, 10, 6, 3, 7, 2} //价值数组
var w = [N]int{0, 4, 6, 2, 2, 5, 1}  //每个物体的重量数组

func main2() {
	for i := 1; i <= n; i++ {
		for j := 1; j <= c; j++ {
			if j > w[i] {
				/*
								2）. j>=w[i] 的情况，这时背包容量可以放下第 i 件物品，我们就要考虑拿这件物品是否能获取更大的价值。
					如果拿取，m[ i ][ j ]=m[ i-1 ][ j-w[ i ] ] + v[ i ]。 这里的m[ i-1 ][ j-w[ i ] ]指的就是考虑了i-1件物品，背包容量为j-w[i]时的最大价值，也是相当于为第i件物品腾出了w[i]的空间。
					如果不拿，m[ i ][ j ] = m[ i-1 ][ j ] , 同（1）
					究竟是拿还是不拿，自然是比较这两种情况那种价值最大。
				*/
				m[i][j] = int(math.Max(float64(m[i-1][j]), float64(m[i-1][j-w[i]]+v[i])))
			} else {
				//背包容量不足加入第i个物品,价值和第[i-1] 是一样的
				m[i][j] = m[i-1][j]
			}
		}
	}

	//打印价值函数
	for i := 1; i <= n; i++ {
		for j := 1; j <= c; j++ {
			fmt.Print(m[i][j], "\t")
		}
		fmt.Println()
	}

	traceback()

	for i := 1; i <= n; i++ {
		fmt.Print(x[i], "\t")
	}

}

//另起一个 x[ ] 数组，x[i]=0表示不拿，x[i]=1表示拿。
//m[n][c]为最优值，如果m[n][c]=m[n-1][c] ,说明有没有第n件物品都一样，则x[n]=0 ;
// 否则 x[n]=1。当x[n]=0时，由x[n-1][c]继续构造最优解；当x[n]=1时，则由x[n-1][c-w[i]]继续构造最优解。以此类推，可构造出所有的最优解。
func traceback() {
	var c1 = 7
	for i := n; i > 1; i-- {
		if m[i][c1] == m[i-1][c1] {
			x[i] = 0
		} else {
			x[i] = 1
			c1 -= w[i]
		}
	}
	x[1] = IfZero(m[1][c1], 1, 0)
}

func IfZero(x int, y int, z int) int {
	if x > 0 {
		return y
	}
	return z
}
