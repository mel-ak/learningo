package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	iterator := true
	for iterator {
		fmt.Print("Enter a paragraph : ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(checkvowel(text))
	}
}

func checkvowel(parag string) int {
	temps := strings.Split(parag, "")
	counter := 0

	for _, temp := range temps {
		temp = strings.ToLower(temp)
		if temp == "a" || temp == "e" || temp == "i" || temp == "o" || temp == "u" {
			counter++
		}
	}

	return counter

}
