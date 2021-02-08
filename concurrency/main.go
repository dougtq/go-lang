package main

import (
	"fmt"
	"time"
)

func main() {
	//var wg sync.WaitGroup
	c := make(chan string)
	go count("cat", c)
	for msg := range c {
		//msg, open := <-c

		// if !open {
		// 	break
		// }
		fmt.Println(msg)
	}
}

func count(thing string, c chan string) {
	for i := 1; i <= 10; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}
