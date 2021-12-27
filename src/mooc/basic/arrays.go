package main

import "fmt"

func modifyArrVal(a [5]int) {
	a[1] = 20
}

func modifyArrRef(a *[5]int) {
	a[1] = 20
}

func main() {
	// 数组不可动态变化问题，一旦声明了，其长度就是固定的。len() 和 cap() 返回结果始终一样
	// 数组是值类型问题，在函数中传递的时候是传递的值，如果传递数组很大，这对内存是很大开销
	// 数组赋值问题，同样类型的数组（长度一样且每个元素类型也一样）才可以相互赋值，反之不可以。
	//一维数组
	var arr_1 [5]int
	fmt.Println(arr_1)

	var arr_2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr_2)

	arr_3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr_3)

	for i, v := range arr_3 {
		fmt.Println(i, v)
	}

	arr_4 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr_4)

	arr_5 := [5]int{0: 3, 1: 5, 4: 6}
	fmt.Println(arr_5)

	//二维数组
	var arr_6 = [3][5]int{{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {3, 4, 5, 6, 7}}
	fmt.Println(arr_6)

	arr_7 := [3][5]int{{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {3, 4, 5, 6, 7}}
	fmt.Println(arr_7)

	arr_8 := [...][5]int{{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {0: 3, 1: 5, 4: 6}}
	fmt.Println(arr_8)

	var arr = [5]int{1, 2, 3, 4, 5}
	modifyArrVal(arr)
	fmt.Println(arr)
	modifyArrRef(&arr)
	fmt.Println(arr)
}
