package utils

import "log"

func MUST[T any](value T, err error) T {
	if err != nil {
		log.Fatal(err)
	}
	return value
}
