package hello

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	ans := make(map[string]int)
	for _, word := range strings.Fields(s) {
		cnt, exist := ans[word]
		if exist {
			ans[word] = cnt + 1
		} else {
			ans[word] = 1
		}
	}
	return ans
}

func main() {
	wc.Test(WordCount)
}
