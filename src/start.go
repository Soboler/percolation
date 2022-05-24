package main

import (
	"fmt"
	"math/rand"
)

const expCount = 1000
const m = 1000
const n = 1000
const s = m*n + 2

func main() {
	rand.Seed(42)

	for i := 0; i < expCount; i++ {
		dsu := createDSU()
		active := make(map[int]struct{})
		active[0] = struct{}{}
		active[s-1] = struct{}{}
		notActive := createNotActive()

		for true {
			var num = rand.Intn(len(notActive))
			var v = notActive[num]
			active[v] = struct{}{}

			var neighbours = getNeighbours(v)
			for j := 0; j < len(neighbours); j++ {
				var neighbour = neighbours[j]
				_, res := active[neighbour]
				if res {
					dsu.UnionSets(v, neighbour)
				}
			}

			if dsu.FindSet(0) == dsu.FindSet(s-1) {
				break
			}
			notActive = append(notActive[0:num], notActive[num+1:]...)
		}

		fmt.Println(len(active) - 2)
	}
}

func createDSU() DSU {
	var parents []int
	var ranks []int

	dsu := DSU{parent: parents, rank: ranks}
	for i := 0; i < s; i++ {
		dsu.MakeSet(i)
	}
	return dsu
}

func createNotActive() []int {
	var res []int
	for i := 1; i < s-1; i++ {
		res = append(res, i)
	}
	return res
}

func getNeighbours(v int) []int {
	var res []int

	var vertNum = (v - 1) / n
	if vertNum == 0 {
		res = append(res, 0)
		res = append(res, v+n)
	} else if vertNum == m-1 {
		res = append(res, s-1)
		res = append(res, v-n)
	} else {
		res = append(res, v+n)
		res = append(res, v-n)
	}

	var gorzNum = (v - 1) % n
	if gorzNum == 0 {
		res = append(res, v+1)
	} else if gorzNum == n-1 {
		res = append(res, v-1)
	} else {
		res = append(res, v+1)
		res = append(res, v-1)
	}

	return res
}
