package main

import "fmt"

func main() {
	m := map[string]string{
		"name":   "aichio",
		"course": "golang",
	}
	m2 := make(map[string]int)
	var m3 map[string]int
	fmt.Println(m, m2, m3)

	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("Getting values")
	courseName := m["course"]
	fmt.Println(courseName)
	courseName = m["cours"]
	fmt.Println(courseName)
	if courseName, ok := m["course"]; ok {
		fmt.Println(courseName)
	} else {
		fmt.Println("key not found")
	}
	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
