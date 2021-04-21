package main

import (
	"fmt"
)

func main()  {
	var scores [3][5]float64
	for i := 0; i < len(scores); i++ {
		for j := 0; j < len(scores[i]); j++ {
			fmt.Printf("input %d %d:", i + 1, j + 1)
			fmt.Scanln(&scores[i][j])
		}
	}
	// fmt.Println("scores=", scores)
	
	total := 0.0
	person := 0.0
	for i := 0; i < len(scores); i++ {
		sum := 0.0
		for j := 0; j < len(scores[i]); j++ {
			sum += scores[i][j]
		}
		person += float64(len(scores[i]))
		total += sum
		fmt.Printf("%d sum=%v avg=%v\n", i, sum, sum / float64(len(scores[i])))
	}
	fmt.Printf("total=%v avg=%v\n", total, total / person)
}