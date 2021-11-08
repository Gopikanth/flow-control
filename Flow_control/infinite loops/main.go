package main

import (
	"bufio"
	"fmt"
	"myapp/do"
	"os"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ch := make(chan string)
	go do.Going(ch)
	fmt.Println("enter the input")
	for i := 0; i < 5; i++ {
		fmt.Println("->")
		input, _ := reader.ReadString('\n')
		ch <- input
		time.Sleep(time.Second)
	}
}
