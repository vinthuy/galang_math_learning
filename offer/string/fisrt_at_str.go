package string

import "math"

//题目：请实现一个函数用来找出字符流中第一个只出现一次的字符。
//举例说明
//　　例如，当从字符流中只读出前两个字符“go”时，第一个只出现一次的字符是‘g’。
//当从该字符流中读出前六个字符“google”时，第一个只出现1次的字符是”l”。
//
//解题思路
//　　字符只能一个接着一个从字符流中读出来。可以定义一个数据容器来保存字符在字符流中的位置。当一个字符第一次从字符流中读出来时，
//把它在字符流中的位置保存到数据容器里。当这个字符再次从字符流中被读出来时，那么它就不是只出现一次的字符，也就可以被忽略了
//。这时把它在数据容器里保存的值更新成一个特殊的值（比如负值）。
//　　为了尽可能高校地解决这个问题，需要在O(1)时间内往容器里插入一个字符，以及更新一个字符对应的值。这个容器可以用哈希表来实现。
//用字符的ASCII码作为哈希表的键值，而把字符对应的位置作为哈希表的值。

type CharStatistics struct {
	vector map[byte]int
}

func (this *CharStatistics) insert(index int, ch byte) {
	if ch > 255 {
		panic("invalid input")
	}
	if this.vector[ch] == -1 {
		this.vector[ch] = index
	} else {
		this.vector[ch] = -2
	}
}

func (this *CharStatistics) Init() {
	this.vector = make(map[byte]int)
	for i := 0; i < 256; i++ {
		this.vector[byte(i)] = -1
	}
}

func FirstFindStr(input string) string {
	if input == "" {
		panic("invalid input")
	}
	var cst = new(CharStatistics)
	cst.Init()
	for i, s := range input {
		cst.insert(i, byte(s))
	}
	var minIndex = math.MaxInt32
	var minCh byte
	for _, s := range input {
		if cst.vector[byte(s)] >= 0 && cst.vector[byte(s)] < minIndex {
			minIndex = cst.vector[byte(s)]
			minCh = byte(s)
		}
	}
	return string(minCh)
}
