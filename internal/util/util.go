package util

import "log"

func ErrorHandler(err error, fatal bool) {
	if fatal {
		log.Fatalln(err.Error())
	} else {
		log.Println(err.Error())
	}
}
