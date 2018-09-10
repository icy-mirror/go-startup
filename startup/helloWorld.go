package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(strconv.Itoa(i) + " BuzzFizz")
		} else if i%3 == 0 {
			fmt.Println(strconv.Itoa(i) + " Fizz")
		} else if i%5 == 0 {
			fmt.Println(strconv.Itoa(i) + " Buzz")
		}
	}
	for s := 1; s <= 100; s++ {
		switch {
		case s%15 == 0:
			fmt.Println(strconv.Itoa(s) + " FUZZBIZZ")
		case s%3 == 0:
			fmt.Println(strconv.Itoa(s) + " FUZZ")
		case s%5 == 0:
			fmt.Println(strconv.Itoa(s) + " BIZZ")
		}
	}
}
