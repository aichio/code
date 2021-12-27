package main

import (
	"fmt"
	"math"
)

var (
	aa = 3
	ss = "sss"
	bb = true
)

func variableZereVaule() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialVaule() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println("%d %d %q\n", a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableScopeGlobal() {
	aa = 10
	ss = "abcd"
	bb = false
	fmt.Println(aa, ss, bb)
}

func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, python, golang, javascript)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func variableScopePart(aa int, ss string, bb bool) {
	aa = 5
	ss = "cde"
	bb = true
	fmt.Println(aa, ss, bb)
}

func main() {
	fmt.Println("Hello world!")
	variableZereVaule()
	variableInitialVaule()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)
	variableScopePart(aa, ss, bb)
	fmt.Println(aa, ss, bb)
	variableScopeGlobal()
	fmt.Println(aa, ss, bb)
	enums()
}
