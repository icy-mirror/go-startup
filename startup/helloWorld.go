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

	fmt.Println("start Prime Judging")
	var primeTable = make([]int, 1)
	primeTable[0] = 2
	var cnt int = 1
	for m := 3; m < maxInt; m++ {
		// m を判定する
		var isPrime bool = true
		for _, mm := range primeTable {
			if mm == 0 {
				break
			}
			fmt.Println(strconv.Itoa(m) + " ÷ " + strconv.Itoa(mm))
			var remainder = m % mm
			if remainder == 0 {
				isPrime = false
			}
			if !isPrime {
				break
			}
		}
		if isPrime {
			primeTable = append(primeTable, m)
			cnt++
		}
	}
	fmt.Println(primeTable)

}
