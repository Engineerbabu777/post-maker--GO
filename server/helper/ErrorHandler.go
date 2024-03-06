package helper

import "log"

func ErrorHandler(err error) {
	log.Fatal(err)
}
