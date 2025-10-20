package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	iterator := true
	for iterator {
		fmt.Print("Enter a number to check primitive : ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		text = strings.Trim(text, "\n")

		number, err := strconv.Atoi(text)
		if err != nil {
			log.Fatal(err)
		}

		if number < 2 {
			fmt.Println("Please enter a number above 1")
			continue
		}
		counter := 0
		for i := 2; i <= number; i++ {
			if number%i == 0 {
				counter++
			}

			if counter > 2 {
				break
			}
		}

		if counter < 2 {
			fmt.Println("Prime")
		} else {
			fmt.Println("Not Prime")
		}
	}

}
