package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		// return a / b
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf(
			"unsupported operation: %s", op)
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// todo 了解 reflect 和 runtime
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with arguments "+
		"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

// 可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func swapVal(a, b int) (int, int) {
	b, a = a, b
	return a, b
}

func swapRef(a, b *int) {
	*a, *b = *b, *a
}

func return_slice() []string {
	slice := []string{"a", "b", "c"}
	return slice
}

func main() {
	test := return_slice()
	fmt.Println(test)
	// fmt.Println(eval(3, 4, "*"))
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	q, r := div(3, 4)
	fmt.Println(q, r)
	fmt.Println(apply(pow, 3, 4))
	fmt.Println(apply(
		func(a, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5, 6))
	a, b := 3, 4
	swapRef(&a, &b)
	fmt.Println(a, b)

	a, b = swapVal(a, b)
	fmt.Println(a, b)
}
