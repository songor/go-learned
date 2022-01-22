// 使用通道实现消息传递
package main

import (
	"fmt"
	"time"
)

func thrower(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
		fmt.Println("Threw >>", i)
	}
}

func catcher(c chan int) {
	for i := 0; i < 5; i++ {
		num := <-c
		fmt.Println("Caught <<", num)
	}
}

func S98() {
	// 有缓冲通道
	c := make(chan int, 3)
	go thrower(c)
	go catcher(c)
	time.Sleep(100 * time.Millisecond)
}
