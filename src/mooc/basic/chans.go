package main

import (
	"fmt"
	"time"
)

func producer(ch chan string) {
	fmt.Println("producer start")
	ch <- "a"
	ch <- "b"
	ch <- "c"
	ch <- "d"
	ch <- "e"
	fmt.Println("producer end")
}

func customer(ch chan string) {
	for {
		msg := <-ch
		fmt.Println(msg)
	}
}

func main() {
	// 声明不带缓冲的通道, 不带缓冲的通道，进和出都会阻塞。
	// 定义的是一个无缓冲的 chan，赋值后就陷入了阻塞。
	// ch1 := make(chan string)

	// 声明带10个缓冲的通道, 带缓冲的通道，进一次长度 +1，出一次长度 -1，如果长度等于缓冲长度时，再进就会阻塞。
	// ch2 := make(chan string, 10)

	// 声明只读通道
	// ch3 := make(<-chan string)

	// 声明只写通道
	// ch4 := make(chan<- string)
	ch := make(chan string, 3)
	fmt.Println("main start")
	ch <- "a"
	go func() {
		val := <-ch // 出 chan
		fmt.Println(val)
	}()
	go func() {
		ch <- "b" // 入 chan
	}()
	go func() {
		val := <-ch // 出 chan
		fmt.Println(val)
	}()
	go producer(ch)
	go customer(ch)
	time.Sleep(1 * time.Second)
	fmt.Println(len(ch))
	fmt.Println("main end")
	close(ch)
}
