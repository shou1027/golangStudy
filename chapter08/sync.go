package main

import (
	"fmt"
	"math/rand"
	"time"
)

const goroutins = 3

func main() {
	c := make(chan int)

	for i := 0; i < goroutins; i++ {
		go func(s chan<- int) {
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)

			fmt.Println("処理完了")

			s <- 0
		}(c)
	}

	for i := 0; i < goroutins; i++ {
		<-c
	}

	fmt.Println("すべて完了")
}
