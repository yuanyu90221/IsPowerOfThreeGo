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
