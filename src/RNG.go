package main

import "math/rand"

type RNG struct {
}

func (r *RNG) create() {
	rand.Seed(42)
}

func (r *RNG) get(n int) int {
	return rand.Intn(n)
}
