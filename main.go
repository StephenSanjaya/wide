package main

import "fmt"

func main() {
	// el := []int{1, 2, 2, 3, 4}
	el := []int{4, 10, 1, 8, 9, 2, 7, 6}

	lenEl := len(el)

	swap := false
	for i := 0; i < lenEl-1; i++ {
		swap = false
		for j := 0; j < lenEl-i-1; j++ {
			if el[j] > el[j+1] {
				temp := el[j]
				el[j] = el[j+1]
				el[j+1] = temp
				swap = true
			}
		}
		if !swap {
			break
		}
	}

	fmt.Println(el)
}
