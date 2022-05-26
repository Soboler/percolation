package main

import (
	"fmt"
)

const expCount = 1000

func main() {
	var random = RNG{}
	random.create()

	for i := 0; i < expCount; i++ {
		var experiment = Experiment{}
		experiment.prepare(random, 1000, 1000)
		var result = experiment.run()

		fmt.Println("=====================")
		fmt.Println(result)
		fmt.Print("PrepareTime:   ")
		fmt.Println(experiment.timers["prepare_End"].Sub(experiment.timers["prepare_Start"]))
		fmt.Print("RunTime:   ")
		fmt.Print(experiment.timers["run_End"].Sub(experiment.timers["run_Start"]))
	}
}
