package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

var quotes = []string{
	"The only way to do great work is to love what you do.",
	"Strive not to be a success, but rather to be of value.",
	"The future belongs to those who believe in the beauty of their dreams.",
	"That which does not kill us makes us stronger.",
	"I have not failed. I've just found 10,000 ways that won't work.",
	"The journey of a thousand miles begins with a single step.",
	"Everything you can imagine is real.",
	"To be yourself in a world that is constantly trying to make you something else is the greatest accomplishment.",
	"Happiness is not something readymade. It comes from your own actions.",
	"Success is not final, failure is not fatal: it is the courage to continue that counts.",
	"You miss 100% of the shots you don't take.",
	"Change the world by being yourself.",
	"Where there is love there is life.",
	"Do what you can, with what you have, where you are.",
	"It is never too late to be what you might have been.",
	"The mind is everything. What you think you become.",
	"We can't help everyone, but everyone can help someone.",
	"Go confidently in the direction of your dreams.",
	"Life is a long lesson in humility.",
	"The best way to predict the future is to create it.",
}

var history = []string{}

func menu() {
	fmt.Println("Welcome to random quote generator")
	fmt.Println("Random Quote Generator")
	fmt.Println("1, Next Quote")
	fmt.Println("2, History")
	fmt.Println("3, Analytics")
	fmt.Println("4, Duplicates")
	fmt.Println("0, Quit")
	fmt.Println("Select what you want to do : ")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	repeator := true

	for repeator {
		menu()

		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		if strings.Contains(text, "0") {
			fmt.Println("Nice to chat to you, Good Bye !!!")
			repeator = false
		} else if strings.Contains(text, "1") {
			randomizer := rand.Intn(len(quotes))
			quote, err := randomQuote(uint(randomizer))
			if err != nil {
				log.Fatal(err)
			}

			cls(2)
			fmt.Println("======================QUOTE==================================================")
			fmt.Println("Q#", quote)
			fmt.Println("======================QUOTE==================================================")
			cls(2)
			history = append(history, quote)
		} else if strings.Contains(text, "2") {
			printHistory()
		} else if strings.Contains(text, "3") {
			quoteAnalytics()
		} else if strings.Contains(text, "4") {
			duplicated()
		}

	}
}

func randomQuote(index uint) (string, error) {
	if len(quotes) < int(index) {
		return "", errors.New("Error getting quote at this index")
	}
	return quotes[index], nil
}

func printHistory() {
	if len(history) < 1 {
		fmt.Println("No history registerd")
		return
	}
	cls(2)
	fmt.Println("======================HISTORY==================================================")
	for _, hist := range history {
		fmt.Println("HQ# ", hist)
	}
	fmt.Println("======================HISTORY==================================================")
	cls(2)
}

func quoteAnalytics() {
	length := len(history)
	if length < 1 {
		cls(2)
		fmt.Println("Nothing to analyze, randomize now")
		cls(2)
		return
	}
	duplicate := len(duplicatedQuotes())

	cls(2)
	fmt.Println("======================ANALYTICS==================================================")
	fmt.Println("Total quote generated : ", length)
	fmt.Println("Total quote duplicated : ", duplicate)
	fmt.Println("======================ANALYTICS==================================================")
	cls(2)
}

func duplicated() {
	dups := duplicatedQuotes()
	if len(dups) < 1 {
		fmt.Println("No history registerd")
		return
	}
	cls(2)
	fmt.Println("======================DUPLICATED==================================================")
	for _, dup := range dups {
		fmt.Println("DHQ# ", dup)
	}
	fmt.Println("======================DUPLICATED==================================================")
	cls(2)
}

func duplicatedQuotes() []string {
	dupls := []string{}
	for _, h := range history {
		dupler := 0
		if in_array(dupls, h) {
			continue
		}
		for _, q := range history {
			if strings.Compare(h, q) == 0 {
				dupler++
			}
			if dupler > 1 {
				if in_array(dupls, h) {
					break
				}
				dupls = append(dupls, h)
			}

		}
	}
	return dupls
}

func in_array(arr []string, value string) bool {
	for _, ar := range arr {
		if strings.Compare(ar, value) == 0 {
			return true
		}
	}
	return false
}

func cls(limit int) {
	for i := 0; i <= limit; i++ {
		fmt.Println("")
	}
}
