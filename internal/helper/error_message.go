package helper

import "strings"

func ValidatorErrMessage(err string) string {
	sep1 := strings.Split(err, ": ")
	sep2 := strings.Split(sep1[1], ".")

	return sep2[0]
}
