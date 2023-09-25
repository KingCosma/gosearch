package gosearch

import (
	"strings"
)

func SearchWords(slice []string, s string) []string {
	strarr := strings.Fields(s)
	var r []string
	for _, x := range slice {
		for _, y := range strarr {
			if strings.Contains(strings.ToLower(x), strings.ToLower(y)) && !slicecontains(r, strings.ToLower(x)) {
				r = append(r, x)
			}
		}
	}
	return r
}

func slicecontains(s []string, ss string) bool {
	for _, x := range s {

		if strings.EqualFold(strings.ToLower(x), strings.ToLower(ss)) {
			return true
		}
	}

	return false
}
