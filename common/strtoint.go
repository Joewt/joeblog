package common

import "strconv"

func StrToUint(s string) uint {
	temp, err  := strconv.Atoi(s)
	if err != nil {
		temp = 0
	}
	id := uint(temp)
	return id
}