package gosearch

import (
	"strings"
)

func SearchWords(slice []string, s string) []string {
	strarr := strings.Fields(s)
	var r []string
	for _, x := range slice {
		for _, y := range strarr {
			if strings.Contains(x, y) && !slicecontains(r, x) {
				r = append(r, x)
			}
		}
	}
	return r
}

func slicecontains(s []string, ss string) bool {
	for _, x := range s {

		if x == ss {
			return true
		}
	}

	return false
}
