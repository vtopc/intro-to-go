package main

import (
	"encoding/json"
	"fmt"

	"github.com/shopspring/decimal"
)

type ShopspringDecimal struct {
	Price decimal.Decimal `json:"price"`
}

func shopspringDecimal() {
	p, _ := decimal.NewFromString("1234567890123456.87")

	v := ShopspringDecimal{
		Price: p,
	}

	b, _ := json.Marshal(v)

	fmt.Println("shopspring/decimal.Decimal:", string(b))
}
