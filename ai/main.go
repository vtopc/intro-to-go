package main

import (
	"ai/ai"
)

var data = ai.Data{
	{Input: 1, Want: 0},
	{Input: 0, Want: 1},
}

func main() {
	network := ai.NewNet(3)
	network.Learn(data)
	network.Test(data)
}
