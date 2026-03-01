package main

import (
	"fmt"
	"strings"
)

var productPrices = map[string]float64{
	"Shirt": 29.99,
	"Pants": 49.99,
	"Hat":   9.99,
	"Shoes": 59.99,
}

func calculateItemPrice(itemName string) (float64, bool) {
	basePrice, found := productPrices[itemName]

	if !found {
		if strings.HasSuffix(itemName, "_SALE") {
			originalItemName := strings.TrimSuffix(itemName, "_SALE")
			basePrice, found = productPrices[originalItemName]

			if found {
				salePrice := basePrice * 0.9
				fmt.Printf("%s are on sale for £%.2f\n", strings.TrimSuffix(itemName, "_SALE"), salePrice)
				return salePrice, true
			}
		}

		fmt.Printf("Item not found: %s\n", itemName)
		return 0.0, false
	}
	return basePrice, true
}

func main() {
	userCart := []string{
		"Shirt",
		"Hat",
		"Shoes_SALE",
	}

	var subtotal float64

	for index, item := range userCart {
		price, found := calculateItemPrice(item)
		if found {
			subtotal += price
			fmt.Printf("%d. %s\n", index+1, item)
		}
	}

	fmt.Printf("Subtotal: £%.2f\n", subtotal)
}
