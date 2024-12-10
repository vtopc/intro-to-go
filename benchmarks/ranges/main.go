package main

type S struct {
	i int
	j [1024]byte
}

func sumElements(input []S) int {
	sum := 0
	for _, elem := range input {
		sum += elem.i
	}

	return sum
}

func sumByIndex(input []S) int {
	sum := 0
	for i := 0; i < len(input); i++ {
		sum += input[i].i
	}

	return sum
}
