package main

import (
	"encoding/json"
	"fmt"
	"math/big"
)

type BigFloat struct {
	Price *big.Float `json:"price"`
}

func bigFloat() {
	v := BigFloat{
		Price: big.NewFloat(1234567890123456.87),
	}

	b, _ := json.Marshal(v)

	fmt.Println("big.Float:", string(b))
}
