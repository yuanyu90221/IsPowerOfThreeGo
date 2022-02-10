# IsPowerOfThreeGo

Write a method to judge whether an input number is power of three

## 我的解法

如同 isPowerOfTwo 一樣

可以知道 如果轉成 base3

假設是 3 的次方數 會符合 10*$ 這樣的模式

```golang
package power

import (
	"regexp"
	"strconv"
)

func IsPowerOfThree(n int64) bool {
	base3Pattern := "^10*$"
	input := strconv.FormatInt(n, 3)
	match, _ := regexp.MatchString(base3Pattern, input)
	return match
}
```