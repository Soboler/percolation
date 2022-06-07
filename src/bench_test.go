package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"testing"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func Benchmark(b *testing.B) {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	var random = RNG{}
	random.create()

	for i := 0; i < expCount; i++ {
		var experiment = Experiment{}
		experiment.prepare(random, 10000, 10000)
		var result = experiment.run()

		fmt.Println("=====================")
		fmt.Println(result)
		fmt.Print("PrepareTime:   ")
		fmt.Println(experiment.timers["prepare_End"].Sub(experiment.timers["prepare_Start"]))
		fmt.Print("RunTime:   ")
		fmt.Print(experiment.timers["run_End"].Sub(experiment.timers["run_Start"]))
	}

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}
