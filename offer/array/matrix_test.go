package array

import (
	"testing"
	"fmt"
	"time"
)

func TestPrintMatrixByCircle(t *testing.T) {
	numbers := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8},
		{22, 23, 24, 25, 26, 27, 28, 9},
		{21, 36, 37, 38, 39, 40, 29, 10},
		{20, 35, 34, 33, 32, 31, 30, 11},
		{19, 18, 17, 16, 15, 14, 13, 12},
	}
	PrintMatrixByCircle(numbers, func(v int) {
		fmt.Print(v, "\t")
	})
}

func Test1(t *testing.T) { //不能在主go route 阻塞，否则会报错
	c := make(chan int) //修改2为1就报错，修改2为3可以正常运行
	//c <- 1
	//c <- 1

	//go add(c)
	go add(c)
	fmt.Println(<-c)

	//fmt.Println(<-c)
	//fmt.Println(<-c)
}
func add(c chan int) {
	c <- 1
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}

func Test3(t *testing.T) {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

func Test4(t *testing.T) {
	c := make(chan int)
	o := make(chan bool)

	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			case <-time.After(5 * time.Second):
				o <- true
			}
		}
	}()

	fmt.Println(<-o)

}
func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func Test5(t *testing.T) {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}

func fibonacci1(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func Test6(t *testing.T) {
	c := make(chan int, 10)
	go fibonacci1(cap(c), c)
	for ch := range c {
		fmt.Println(ch)
	}
}

func Test7(t *testing.T) {
	ch := make(chan string)
	close(ch)
	i,ok := <- ch
	if ok{
		fmt.Println(i)
	}else {
		fmt.Println("closed")
	}

}