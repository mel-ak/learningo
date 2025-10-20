package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	iterator := true

	for iterator {

		mprint("Enter a number (q to close) : ")
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		numi := strings.Trim(input, "\n")

		if numi == "q" {
			iterator = false
			mprint("Good Byee !!!")
			break
		} else {

			num, err := strconv.Atoi(numi)
			if err != nil {
				log.Fatal(err)
			}
			mprint(num)

			if num%3 == 0 && num%5 == 0 {
				mprint("FizzBuzz")
			} else if num%3 == 0 {
				mprint("Fizz")
			} else if num%5 == 0 {
				mprint("Buzz")
			} else {
				mprint(num)
			}
		}

	}

}

func mprint(p ...any) {
	fmt.Println(p...)
}

func mread(r io.Reader) {

}
