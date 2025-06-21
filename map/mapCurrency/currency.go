package main

import (
	"fmt"
)

func main() {
	currncy := make(map[string]int)
	currncy["EURO"] = 100
	currncy["USD"] = 80
	currncy["ROUBLE"] = 20
	rangeOverMap(currncy)
	rangeOverMap(deleteElementMap(currncy, "USD"))
}

func rangeOverMap(currency map[string]int) {
	fmt.Printf("All Memeber Are %d\n", len(currency))
	for key, element := range currency {
		fmt.Println(key, element)
	}
}

func deleteElementMap(currency map[string]int, key string) map[string]int {
	delete(currency, key)
	return currency
}
