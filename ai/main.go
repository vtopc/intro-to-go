package main

import (
	"ai/ai"
)

var learnData = ai.Data{
	{Input: 1, Want: 0},
	{Input: 0, Want: 1},
}

var testData = ai.Data{
	{Input: 1, Want: 0},
	{Input: 0, Want: 1},
}

func main() {
	network := ai.NewNet(3)
	network.Learn(learnData)
	network.Test(testData)
}
