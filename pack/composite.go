package pack

import (
	"fmt"
	"unsafe"
)

var a = []int{1, 2, 3, 4}

const N1 = 3

func main1() {
	aa := 1
	fmt.Printf("suba entry address:%p \n", aa)
	fmt.Println(aa)
	aaa(&aa)
	fmt.Println(aa)
	//suba := make([]int, 0, 2)
	//suba = append(suba, 1)
	//fmt.Printf("suba  address:%p %v ,length:%d,cap:%d \n", &suba, suba,len(suba),cap(suba))
	//
	//suba = append(suba, 2)
	//fmt.Printf("suba  address:%p %v ,length:%d,cap:%d \n", &suba, suba,len(suba),cap(suba))
	//
	//suba = append(suba, 3)
	//fmt.Printf("suba  address:%p %v ,length:%d,cap:%d \n", &suba, suba,len(suba),cap(suba))

	//composite(a, &suba, 3, 0)
	//fmt.Println("***********")
	//FullSort(a[0:], 0)
}

func aaa(a *int) {
	fmt.Printf("suba joint address:%p \n", &a)
	*a = 2
}

type Slice struct {
	ptr unsafe.Pointer // Array pointer
	len int            // slice length
	cap int            // slice capacity
}

//每一次从集合中选出一个元素，然后对剩余的集合（n-1）进行一次k-1组合。
//组合函数
func composite(a []int, suba1 *[]int, k int, start int) {
	suba := *suba1
	fmt.Printf("suba joint address:%p %v ,length:%d,cap:%d \n", &suba, suba, len(suba), cap(suba))

	if k == 0 {
		fmt.Println(suba)
		return
	}
	var m = len(a)

	for i := start; i <= m-k; i++ {
		suba = append(suba, a[i])
		/*slice1 := (*Slice)(unsafe.Pointer(&suba))
		fmt.Println(slice1.ptr)*/
		if k > 0 {
			fmt.Printf("suba  entery address:%p %v ,length:%d,cap:%d \n", &suba, suba, len(suba), cap(suba))
			composite(a, &suba, k-1, i+1)
			if len(suba) > 0 {
				suba = suba[0 : len(suba)-1]
				//fmt.Printf("suba  remove address:%p %v,length:%d,cap:%d \n", &suba, suba,len(suba),cap(suba))
			}
		}
	}
}

func swap(o []int, i int, j int) {
	tmp := o[i]
	o[i] = o[j]
	o[j] = tmp
}

func FullSort(a []int, start int) {
	if start == len(a)-1 {
		fmt.Println(a)
		return
	}

	for i := start; i < len(a); i++ {
		if i > start {
			swap(a, start, i)
		}
		FullSort(a, start+1)
		if i > start {
			swap(a, i, start)
		}
	}
}
