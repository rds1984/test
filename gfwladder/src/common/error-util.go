package common

import "log"

func HandleError(msg string, err error) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
