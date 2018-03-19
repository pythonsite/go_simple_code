package main

func longestCommonPrefix(strs []string) string {
	row := len(strs)
	if row == 0 {
		return ""
	}
	if row == 1 {
		return strs[0]
	}
	lcp := ""
	minLen := (1 << 30)
	// 计算最长
	for _, v := range strs {
		vLen := len(v)
		if minLen > vLen {
			minLen = vLen
		}
	}
BREAKPOINT:
	for i := 0; i < minLen; i++ {
		for j := 0; j < row; j++ {
			if strs[j][i] != strs[0][i] {
				break BREAKPOINT
			}
		}
		lcp += strs[0][i : i+1]
	}
	return lcp
}

func main() {
}
