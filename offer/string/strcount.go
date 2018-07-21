package string

import "fmt"

//题目：在字符串中找出第一个只出现一次的字符。
//解题思路：
//第一种：直接求解：
//
//　　从头开始扫描这个字符串中的每个字符。当访问到某字符时拿这个字符和后面的每个字符相比较，如果在后面没有发现重复的字符，则该字符就是只出现一次的字符。
// 如果字符串有n个字符，每个字符可能与后面的O(n）个字符相比较，因此这种思路的时间复杂度是O(n^2)。
//
//第二种：记录法
//
//　　由于题目与字符出现的次数相关， 我们是不是可以统计每个字符在该字符串中出现的次数？要达到这个目的，我们需要一个数据容器来存放每个字符的出现次数。
// 在这个数据容器中可以根据字符来查找它出现的次数，也就是说这个容器的作用是把一个字符映射成二个数字。在常用的数据容器中， 哈希表正是这个用途。
//　　为了解决这个问题，我们可以定义哈希表的键（Key）是字符， 而值(Value ）是该字符出现的次数。同时我们还需要从头开始扫描字符串两次。
// 第一次扫描字符串时，每扫描到一个字符就在哈希表的对应项中把次数加1 。接下来第二次扫描时， 每扫描到一个字符就能从哈希表中得到该字符出现的次数。
// 这样第一个只出现一次的字符就是符合要求的输出。
//　　第一次扫描时，在哈希表中更新一个字符出现的次数的时间是O(n) 。
// 如果字符串长度为n， 那么第一次扫描的时间复杂度是O(n)。第二次扫描时，同样0(1)能读出一个字符出现的次数，所以时间复杂度仍然是O(n)。这样算起来，总的时间复杂度是O(n)。

func FirstCountStr(str string) string {
	if str == "" || len(str) == 0 {
		return ""
	}
	charCountMap := make(map[byte]int, 256)
	for _, s := range str {
		charCountMap[byte(s)]++
	}
	fmt.Println(charCountMap)
	for _, s := range str {
		bys := byte(s)
		if charCountMap[bys] == 1 {
			return string(s)
		}
	}
	return ""
}

//题目一：输入一个英文句子，翻转句子中单词的顺序，但单词内字啊的顺序不变。为简单起见，标点符号和普通字母一样处理。
//举例说明
//例如输入字符串”I am a student. ”，则输出”student. a am I”。

//解题思路
// 第一步翻转句子中所有的字符。比如翻转“I am a student. ”中所有的字符得到”.tneduts a m a I”，此时不但翻转了句子中单词的顺序，
// 连单词内的字符顺序也被翻转了。第二步再翻转每个单词中字符的顺序，就得到了”student. a am I”。这正是符合题目要求的输出。
func ReversByte(str []byte, start, end int) []byte {

	if start < 0 || end < 0 || start >= len(str) || end >= len(str) || start > end {
		panic("invalid input!!")
	}
	for start < end {
		tmp := str[start]
		str[start] = str[end]
		str[end] = tmp
		start++
		end --
	}
	return str
}

//student a am I
func RerverString(str string) string {
	var srcbyte = []byte(str)
	//1.
	srcbyte = ReversByte(srcbyte, 0, len(srcbyte)-1)
	//2
	start, end := 0, 0
	for start < len(srcbyte) {
		if srcbyte[start] == ' ' {
			start++
			end++
		} else if end == len(srcbyte) || srcbyte[end] == ' ' {
			ReversByte(srcbyte, start, end-1)
			end++
			start = end
		} else {
			end++
		}
	}
	return string(srcbyte)
}

//题目二：字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。请定义一个函数实现字符串左旋转操作的功能。
//举例说明
//比如输入字符串”abcefg”和数字2，该函数将返回左旋转2 位得到的结”cdefab”。
//
//解题思路
//　　以”abcdefg”为例，我们可以把它分为两部分。由于想把它的前两个字符移到后面，
// 我们就把前两个字符分到第一部分，把后面的所有字符都分到第二部分。
// 我们先分别翻转这两部分，于是就得到”bagfedc”。接下来我们再翻转整个字符串，
// 得到的”cde也ab”同lj好就是把原始字符串左旋转2 位的结果。

func RonateWorld(str string, splintIndex int) string {
	var srcbyte = []byte(str)
	ReversByte(srcbyte, 0, splintIndex-1)
	ReversByte(srcbyte, splintIndex, len(srcbyte)-1)
	ReversByte(srcbyte, 0, len(srcbyte)-1)
	return string(srcbyte)
}

