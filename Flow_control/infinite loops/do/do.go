package do

import "log"

func Going(ch chan string) {
	for {
		msg := <-ch
		log.Println(msg)
	}
}
