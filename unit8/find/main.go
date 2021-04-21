package main

import (
	"fmt"
)

func main()  {
	names := [5]string{"xxx", "yyy", "zzz", "sss", "mmm"}
	var heroName = ""
	fmt.Println("input please:")
	fmt.Scanln(&heroName)

	// 顺序查找一
	for i := 0; i < len(names); i++ {
		if heroName == names[i] {
			fmt.Printf("get...%v index=%v\n", heroName, i)
			break
		} else if i == len(names) - 1 {
			fmt.Printf("no...%v\n", heroName)
		} 
	}

	// 顺序查找二
	index := -1
	for i := 0; i < len(names); i++ {
		if heroName == names[i] {
			index = i
			break
		}
	}
	if index != -1 {
		fmt.Printf("get...%v index=%v\n", heroName, index)
	} else {
		fmt.Printf("no...%v\n", heroName)
	}
}