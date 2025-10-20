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

	fmt.Print("Enter string to check Palindrome : ")
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fake := reverser(strings.Trim(text, "\n"))
	if strings.Trim(text, "\n") == fake {
		fmt.Println("It is a palindrome")
	} else {
		fmt.Println("It is not a palindrome")
	}

}

func reverser(str string) string {
	temp := ""
	strray := strings.Split(str, "")

	for i := len(strray) - 1; i >= 0; i-- {
		temp += strray[i]
	}

	return strings.Trim(temp, "\n")
}
