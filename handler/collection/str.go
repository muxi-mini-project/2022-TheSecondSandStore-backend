package collection

import (
	"strconv"
	"strings"
)

func IntSliceToString(s []int) string {

	res := ""
	for _, v := range s {
		i := strconv.Itoa(v)
		res = res + i + " "
	}
	res = strings.TrimRight(res, " ")
	return res
}

func StringToStringSlice(str string) []string {

	return strings.Split(str, " ")
}
