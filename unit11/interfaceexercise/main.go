package main

import (
	"fmt"
	"sort"
	"math/rand"
)

type Hero struct {
	Name string
	Age int
} 

type HeroSlice []Hero

func (heroSlice HeroSlice) Len() int {
	return len(heroSlice)
}

func (heroSlice HeroSlice) Less(i, j int) bool {
	// return heroSlice[i].Age < heroSlice[j].Age
	return heroSlice[i].Name < heroSlice[j].Name
}

func (heroSlice HeroSlice) Swap(i, j int) {
	// temp := heroSlice[i]
	// heroSlice[i] = heroSlice[j]
	// heroSlice[j] = temp

	heroSlice[i], heroSlice[j] = heroSlice[j], heroSlice[i]
}

func main()  {
	// 定义一个切片
	var intSlice = []int{0, -1, 10, 7, 90}
	// 对切片进行排序
	sort.Ints(intSlice)
	fmt.Println(intSlice) // [-1 0 7 10 90]

	// 对结构体切片进行排序
	var hs HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero {
			Name : fmt.Sprintf("英雄~%d", rand.Intn(100)),
			Age : rand.Intn(100),
		}
		hs = append(hs, hero)
	}
	fmt.Println("sort before...")
	for _, val := range hs {
		fmt.Println(val)
	}
	sort.Sort(hs)
	fmt.Println("sort after...")
	for _, val := range hs {
		fmt.Println(val)
	}
}