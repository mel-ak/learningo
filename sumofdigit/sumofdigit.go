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

		fmt.Print("Enter a string number to be added ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		sum := 0
		for _, val := range strings.Split(text, "") {
			convable := strings.Trim(val, "\n")

			if convable == "" {
				continue
			}
			temp, err := strconv.Atoi(convable)
			if err != nil {
				log.Fatal(err)
			}

			sum += temp
		}
		fmt.Println(sum)
	}

}
