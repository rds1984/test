package common

import "log"

func HandleError(err error, msg ...interface{}) {
	if err != nil {
		switch len(msg) {
		case 0:
			log.Fatalln(err)
			break
		case 1:
			log.Fatalln(msg[0], err)
			break
		case 2:
			log.Fatalln(msg[0], msg[1], err)
			break
		case 3:
			log.Fatalln(msg[0], msg[1], msg[2], err)
			break
		}
	}
}

func HandleInfo(err error, msg ...interface{}) {
	switch len(msg) {
	case 0:
		log.Println(err)
		break
	case 1:
		log.Println(msg[0], err)
		break
	case 2:
		log.Println(msg[0], msg[1], err)
		break
	case 3:
		log.Println(msg[0], msg[1], msg[2], err)
		break
	}
}
