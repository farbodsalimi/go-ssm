package utils

import (
	"log"
)

// ErrorExit bails out on error
func ErrorExit(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
