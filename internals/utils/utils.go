package utils

import "log"

func MUST[T any](value T, err error, msg string) T {
	if err != nil {
		log.Fatal(msg, " : ", err)
	}
	return value
}
