package main

import "fmt"

const expCount = 100
const n = 1000
const m = 1000

func main() {
	var random = RNG{}
	random.create()

	var count int64
	for i := 0; i < expCount; i++ {
		var experiment = Experiment{}
		experiment.prepare(random, m, n)
		experiment.run()

		count += int64(experiment.step)
		//var result = experiment.run()
		//fmt.Println(result)
	}
	fmt.Println(float64(count) / float64(m*n*expCount))
}
