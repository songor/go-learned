// 使用通道同步 goroutine
package main

import (
	"fmt"
	"time"
)

func printNumbers3(w chan bool)  {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d ", i)
	}
	w <- true
}

func printLetters3(w chan bool)  {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}
	w <- true
}

func S97() {
	w1, w2 := make(chan bool), make(chan bool)
	go printNumbers3(w1)
	go printLetters3(w2)
	<- w1
	<- w2
}