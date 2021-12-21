package main

import "fmt"

func main() {
	// 切片是一种动态数组，比数组操作灵活，长度不是固定的，可以进行追加和删除。 len() 和 cap() 返回结果可相同和不同

	var sli_1 []int //nil 切片
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_1), cap(sli_1), sli_1)

	var sli_2 = []int{} //空切片
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_1), cap(sli_2), sli_2)

	var sli_3 = []int{1, 2, 3, 4, 5}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_3), cap(sli_3), sli_3)

	sli_4 := []int{1, 2, 3, 4, 5}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_4), cap(sli_4), sli_4)

	var sli_5 []int = make([]int, 5, 8)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_5), cap(sli_5), sli_5)

	sli_6 := make([]int, 5, 9)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_6), cap(sli_6), sli_6)

	sli := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli), cap(sli), sli)

	fmt.Println("sli[1] ==", sli[1])
	fmt.Println("sli[:] ==", sli[:])
	fmt.Println("sli[1:] ==", sli[1:])
	fmt.Println("sli[:4] ==", sli[:4])

	fmt.Println("sli[0:3] ==", sli[0:3])
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli[0:3]), cap(sli[0:3]), sli[0:3])

	fmt.Println("sli[0:3:4] ==", sli[0:3:4])
	// index 0:3 cap=4
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli[0:3:4]), cap(sli[0:3:4]), sli[0:3:4])
}
