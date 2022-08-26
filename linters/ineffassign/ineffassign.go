package ineffassign

func storeCounter(counter int) {
	_ = counter
}

func ineffassign() {
	var counter int

	// do something

	counter++
	storeCounter(counter)

	// do something else
	counter++
}
