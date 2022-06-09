package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"strconv"
)

type RNG struct {
	reader io.Reader
}

func (r *RNG) create() {
}

func (r *RNG) get(n int) int {
	var a = big.NewInt(int64(n))
	bigInt, err := rand.Int(rand.Reader, a)
	if err != nil {
		fmt.Println("error:", err)
		r.get(n)
	}

	res, err := strconv.Atoi(bigInt.String())
	if err != nil {
		fmt.Println("error:", err)
	}

	return res
}
