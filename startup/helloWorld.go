package main

import (
	"fmt"
	"strconv"
)

const maxInt int = 1000

func main() {
	var arr [maxInt]int = [maxInt]int{}
	for n := 0; n < maxInt; n++ {
		fmt.Println(strconv.Itoa(n))
		arr[n] = n + 1
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
