package main

import (
	"encoding/json"
	"fmt"
)

type Float struct {
	Price float64 `json:"price"`
}

func float() {
	v := Float{
		Price: 1234567890123456.87,
	}

	b, _ := json.Marshal(v)

	fmt.Println("float64:", string(b))
}
