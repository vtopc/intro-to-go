package main

type S struct {
	i int
	j [1024]byte
}

func sumValueElements(input []S) int {
	sum := 0
	for _, elem := range input {
		sum += elem.i
	}

	return sum
}

func sumValuesByIndex(input []S) int {
	sum := 0
	for i := 0; i < len(input); i++ {
		sum += input[i].i
	}

	return sum
}

func sumPointerElements(input []*S) int {
	sum := 0
	for _, elem := range input {
		sum += elem.i
	}

	return sum
}

func sumPointersByIndex(input []*S) int {
	sum := 0
	for i := 0; i < len(input); i++ {
		sum += input[i].i
	}

	return sum
}
