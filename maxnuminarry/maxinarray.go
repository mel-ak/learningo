package main

import "fmt"

func main() {
	numarr := []int{1, 8, 4, 5, 6}

	max := numarr[0]
	for _, num := range numarr {
		if num > max {
			max = num
		}
	}

	fmt.Println(max)

}
