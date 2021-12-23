package main

import (
	"fmt"
	"time"
)

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func t1() int {
	a := 1
	defer func() {
		a++
	}()
	return a
}

func t2() (a int) {
	defer func() {
		a++
	}()
	return 1
}

func t3() (b int) {
	a := 1
	defer func() {
		a++
	}()
	return 1
}

func t4() (a int) {
	defer func(a int) {
		a++
	}(a)
	return 1
}

func GoA() {
	defer (func() {
		if err := recover(); err != nil {
			fmt.Println("panic:" + fmt.Sprintf("%s", err))
		}
	})()
	// panic("error")
	go GoB()
}

func GoB() {
	panic("error")
}

func main() {
	defer fmt.Println("1")
	// 结论：当`os.Exit()`方法退出程序时，defer不会被执行。
	// os.Exit(0)
	fmt.Println("main")
	x := 1
	y := 2
	defer calc("A", x, calc("B", x, y))
	x = 3
	defer calc("C", x, calc("D", x, y))
	y = 4
	fmt.Println("===================")
	a := t1()
	fmt.Println("t1", a)
	a = t2()
	fmt.Println("t2", a)
	a = t3()
	fmt.Println("t3", a)
	a = t4()
	fmt.Println("t4", a)

	// `GoB()` panic 捕获不到。defer 只对当前协程有效。
	GoA()
	time.Sleep(1 * time.Second)
	fmt.Println("main")
}
