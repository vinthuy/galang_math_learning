package strreplace

//题目：请实现一个函数，把字符串中的每个空格替换成"%20"，例如“We are happy.”，则输出“We%20are%20happy.”。
// 思路:
// 1.求空格大小
// 2.新字符总的大小， 遍历字符串，当遇到空格时，写入替换字符，否则，把遇到字符写入到新字符串中
func ReplaceBlank(str string, c byte, destStr string) string {
	oldLength := len(str)
	blackCount := 0
	for _, s := range str {
		if byte(s) == c {
			blackCount++
		}
	}
	newLength := oldLength + blackCount*(len(destStr)-1)
	res := make([]byte, newLength, newLength)
	j := newLength - 1
	for i := oldLength - 1; i >= 0; i-- {
		if str[i] == c {
			for m := len(destStr) - 1; m >= 0; m-- {
				res[j] = byte(destStr[m])
				j--
			}
		} else {
			res[j] = byte(str[i])
			j--
		}
	}
	return string(res)
}
