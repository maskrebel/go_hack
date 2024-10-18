package main

import "fmt"

func priceCheck(products []string, productPrice []float64, productSold []string, soldPrice []float64) int32 {
	mapProductPrice := make(map[string]float64)
	for i, product := range products {
		mapProductPrice[product] = productPrice[i]
	}

	res := int32(0)

	for i, val := range productSold {
		if mapProductPrice[val] != soldPrice[i] {
			res++
		}
	}

	return res
}

func main() {
	products := []string{"eggs", "milk", "cheese"}
	productsPrices := []float64{2.89, 3.29, 5.79}
	productsSold := []string{"eggs", "eggs", "cheese", "milk"}
	soldPrice := []float64{2.89, 2.99, 5.97, 3.29}

	fmt.Println(priceCheck(products, productsPrices, productsSold, soldPrice))

}
