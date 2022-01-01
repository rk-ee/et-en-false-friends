package pkg

import (
	"strconv"
	"strings"
)

func StringsHasBothix(s string, bothix string) bool {
	return strings.HasPrefix(s, bothix) && strings.HasSuffix(s, bothix)
}
func StringsTrimBothix(s string, bothix string) string {
	return strings.TrimPrefix(strings.TrimSuffix(s, bothix), bothix)
}
func StrconvSliceAtoi(ss []string) (out []int, err error) {
	for _, s := range ss {
		var i int
		i, err = strconv.Atoi(s)
		if err != nil {
			break
		}
		out = append(out, i)
	}
	return
}
