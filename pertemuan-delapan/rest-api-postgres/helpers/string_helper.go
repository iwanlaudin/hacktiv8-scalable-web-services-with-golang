package helpers

import "strconv"

func Str2Int(s string) (int, error) {
	return strconv.Atoi(s)
}
