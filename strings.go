package goutils

import (
	"strconv"
	"strings"
)

func GetIntArrFromString(s string) []int {
	if strings.TrimSpace(s) == "" {
		return nil
	}

	out := make([]int, 0)

	for _, v := range strings.Split(s, ",") {
		n, err := strconv.Atoi(strings.TrimSpace(v))
		if err == nil {
			out = append(out, n)
		}
	}

	if len(out) == 0 {
		return nil
	}

	return out
}
