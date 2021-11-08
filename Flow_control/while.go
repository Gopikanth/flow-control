package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	i := 500
	for i > 70 {
		i = rand.Intn(500) + 1
		fmt.Println("i is", i)
		if i > 70 {
			fmt.Println("the loops continues", i)
		} else {
			fmt.Println("loop ends", i)
		}
	}

}
