package string

//题目：请实现一个函数用来匹配包含‘.’和‘*’的正则表达式。模式中的字符’.’表示任意一个字符，而‘*’表示它前面的字符可以出现任意次（含0次）。
//本题中，匹配是指字符串的所有字符匹配整个模式。例如，字符串“aaa”与模式“a.a”和“ab*ac*a”匹配，但与“aa.a”及“ab*a”均不匹配。

/**
题目解析
　　每次从字符串里拿出一个字符和模式中的字符去匹配。先来分析如何匹配一个字符。如果模式中的字符ch是‘.’，那么它可以匹配字符串中的任意字符。
如果模式中的字符ch不是’.’而且字符串中的字符也是ch，那么他们相互匹配。当字符串中的字符和模式中的字符相匹配时，接着匹配后面的字符。
　　相对而言当模式中的第二个字符不是‘*’时问题要简单很多。如果字符串中的第一个字符和模式中的第一个字符相匹配，
那么在字符串和模式上都向后移动一个字符，然后匹配剩余的字符串和模式。如果字符串中的第一个字符和模式中的第一个字符不相匹配，则直接返回false。
　　当模式中的第二个字符是‘*’时问题要复杂一些，因为可能有多种不同的匹配方式。一个选择是在模式上向后移动两个字符。
这相当于‘*’和它面前的字符被忽略掉了，因为‘*’可以匹配字符串中0个字符。如果模式中的第一个字符和字符串中的第一个字符相匹配时，则在字符串向后移动一个字符
，而在模式上有两个选择：我们可以在模式上向后移动两个字符，也可以保持模式不变。
 */

func MatchStr(input, regx string) bool {
	if input == "" || regx == "" {
		return false
	}
	return matchCore(input, regx, 0, 0)
}
func matchCore(input, regx string, i, p int) bool {

	if i >= len(input) && p >= len(regx) {
		return true
	}
	if i != len(input) && p >= len(regx) {
		return false
	}

	if p+1 < len(regx) && regx[p+1] == '*' {
		if i >= len(input) {
			return matchCore(input, regx, i, p+2)
		} else {
			if regx[p+1] == input[i+1] || regx[p] == '.' {
				// 匹配串向后移动一个位置，模式串向后移动两个位置
				// 匹配串向后移动一个位置，模式串不移动
				// 匹配串不移动，模式串向后移动两个位置
				return matchCore(input, regx, i+1, p+2) || matchCore(input, regx, i+1, p) || matchCore(input, regx, i, p+2)
			} else {
				return matchCore(input, regx, i, p+2)
			}
		}
	}

	if i > len(input) {
		return false
	}

	if regx[p] == input[i] || regx[p] == '.' {
		return matchCore(input, regx, i+1, p+1)
	}
	return false
}

