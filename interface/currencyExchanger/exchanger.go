package main

import "fmt"

type currencyExchnageUnit struct {
	EUR    float64
	USD    float64
	ROUBLE float64
}

func (c currencyExchnageUnit) rialToEUR(rial float64) float64 {
	euro := float64(rial / c.EUR)
	return euro
}

func (c currencyExchnageUnit) rialToUSD(rial float64) float64 {
	usd := float64(rial / c.USD)
	return usd
}

func (c currencyExchnageUnit) rialToROUBLE(rial float64) float64 {
	rouble := float64(rial / c.ROUBLE)
	return rouble
}

type exchangerInterface interface {
	rialToROUBLE(rial float64) float64
	rialToUSD(rial float64) float64
	rialToEUR(rial float64) float64
}

func main() {
	myMony := currencyExchnageUnit{EUR: 100000, USD: 80000, ROUBLE: 20000}
	fmt.Println(myMony)
	exchenger(myMony)
}

func exchenger(e exchangerInterface) {
	rial := 10000000
	usd := e.rialToUSD(float64(rial))
	euro := e.rialToEUR(float64(rial))
	rouble := e.rialToROUBLE(float64(rial))
	fmt.Println(usd, euro, rouble)
}
