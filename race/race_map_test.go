package main

import "testing"

func Test_race(t *testing.T) {
	x := make(T)
	race(x)
}

func Test_raceWithRecover(t *testing.T) {
	x := make(T)
	raceWithRecover(x)
}
