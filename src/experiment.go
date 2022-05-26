package main

import (
	"fmt"
	"time"
)

type Experiment struct {
	random RNG
	timers map[string]time.Time
	m      int
	n      int
	s      int

	dsu       DSU
	active    map[int]struct{}
	notActive []int
	step      int
}

func (exp *Experiment) prepare(r RNG, m int, n int) {
	exp.timers = make(map[string]time.Time)
	exp.random = r
	exp.n = n
	exp.m = m
	exp.s = exp.m*exp.n + 2
	exp.step = 0

	exp.timers["prepare_Start"] = time.Now()
	exp.createActive()
	exp.createNotActive()
	exp.createDSU()
	exp.timers["prepare_End"] = time.Now()
}

func (exp *Experiment) createActive() {
	exp.active = make(map[int]struct{})
	exp.active[0] = struct{}{}
	exp.active[exp.s-1] = struct{}{}
}

func (exp *Experiment) createNotActive() {
	for i := 1; i < exp.s-1; i++ {
		exp.notActive = append(exp.notActive, i)
	}
}

func (exp *Experiment) createDSU() {
	var parents []int
	var ranks []int

	exp.dsu = DSU{parent: parents, rank: ranks}
	for i := 0; i < exp.s; i++ {
		exp.dsu.MakeSet(i)
	}
}

func (exp *Experiment) run() float64 {
	exp.timers["run_Start"] = time.Now()

	for true {
		var num = exp.random.get(exp.s - 2 - exp.step)
		var v = exp.notActive[num]
		exp.step += 1
		exp.active[v] = struct{}{}

		var neighbours = exp.getNeighbours(v)
		for i := 0; i < len(neighbours); i++ {
			var neighbour = neighbours[i]
			_, res := exp.active[neighbour]
			if res {
				exp.dsu.UnionSets(v, neighbour)
			}
		}

		if exp.dsu.FindSet(0) == exp.dsu.FindSet(exp.s-1) {
			break
		}
		exp.notActive = append(exp.notActive[0:num], exp.notActive[num+1:]...)
	}

	fmt.Println(len(exp.active) - 2)

	exp.timers["run_End"] = time.Now()
	return float64(len(exp.active)-2) / float64(exp.s-2)
}

func (exp *Experiment) getNeighbours(v int) []int {
	var res []int

	var vertNum = (v - 1) / exp.n
	if vertNum == 0 {
		res = append(res, 0)
		res = append(res, v+exp.n)
	} else if vertNum == exp.m-1 {
		res = append(res, exp.s-1)
		res = append(res, v-exp.n)
	} else {
		res = append(res, v+exp.n)
		res = append(res, v-exp.n)
	}

	var gorzNum = (v - 1) % exp.n
	if gorzNum == 0 {
		res = append(res, v+1)
	} else if gorzNum == exp.n-1 {
		res = append(res, v-1)
	} else {
		res = append(res, v+1)
		res = append(res, v-1)
	}

	return res
}
