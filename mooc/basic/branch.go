package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf(
			"Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	const filename = "README.md"
	// if 语句可以赋初值
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	// fmt.Println(contents) 变量在 if 生成器中产生, if 外无法获取该值
	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(82),
		grade(91),
		grade(100),
		// grade(101),
		// grade(-3),
	)
}
