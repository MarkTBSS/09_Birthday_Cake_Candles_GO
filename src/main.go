package main

import (
	"fmt"
	"sort"
)

func main() {
	var candles []int32
	candles = []int32{4, 4, 4, 1, 3}
	fmt.Println("Original Integers ============ :", candles)
	sort.Slice(candles, func(i, j int) bool {
		return candles[i] > candles[j]
	})
	fmt.Println("Sorted Integers ============== :", candles)

	var countEqualElement int32
	countEqualElement = 1
	for i := 0; i < len(candles)-1; i++ {
		if candles[i] == candles[i+1] {
			countEqualElement++
		}
	}
	fmt.Println("Amount of Highest Value ====== :", countEqualElement)
}
