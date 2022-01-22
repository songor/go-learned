// goroutine 使用示例
package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		//fmt.Printf("%d ", i)
	}
}

func printLetters() {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		//fmt.Printf("%c ", i)
	}
}

// 使用等待组
func printNumbers2(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d ", i)
	}
	wg.Done()
}

func printLetters2(wg *sync.WaitGroup) {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}
	wg.Done()
}

func print() {
	printNumbers()
	printLetters()
}

func goPrint() {
	go printNumbers()
	go printLetters()
}

func S96() {
	var wg sync.WaitGroup
	wg.Add(2)
	go printNumbers2(&wg)
	go printLetters2(&wg)
	wg.Wait()
}
