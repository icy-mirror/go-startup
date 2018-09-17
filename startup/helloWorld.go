package main

import (
	"fmt"
	"strconv"
)

const maxInt int = 100

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
	var arr [maxInt]int = [maxInt]int{}
	for n := 0; n < maxInt; n++ {
		fmt.Println(strconv.Itoa(n))
		arr[n] = n + 1
	}

	for s := 0; s < len(arr); s++ {
		fmt.Println(strconv.Itoa(arr[s]) + " ")
		switch {
		case arr[s]%15 == 0:
			fmt.Println(strconv.Itoa(arr[s]) + " FUZZBIZZ")
		case arr[s]%3 == 0:
			fmt.Println(strconv.Itoa(arr[s]) + " FUZZ")
		case arr[s]%5 == 0:
			fmt.Println(strconv.Itoa(arr[s]) + " BIZZ")
		}
	}

	fmt.Println("start prime Judging")
	var primeTable [maxInt]int = [maxInt]int{}
	primeTable[0] = 2
	var cnt int = 1
	for m := 3; m < maxInt; m++ {
		// m を判定する
		var isPrime bool = true
		for mm := 0; mm < len(primeTable); mm++ {
			if primeTable[mm] == 0 {
				break
			}
			fmt.Println(strconv.Itoa(m) + " ÷ " + strconv.Itoa(primeTable[mm]))
			var remainder = m % primeTable[mm]
			if remainder == 0 {
				isPrime = false
			}
			if !isPrime {
				break
			}
		}
		if isPrime {
			primeTable[cnt] = m
			cnt++
		}
	}
	fmt.Println(primeTable)

}
