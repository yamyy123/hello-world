package hello

import "log"

func Add(x, y int) int {
	log.Println(x + y)
	log.Println("hi")
	return x + y
}
