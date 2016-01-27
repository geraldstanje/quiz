package main

import (
	"flag"
	"fmt"
	"lcwsolver"
	"log"
	"os"
	"time"
)

var flagFilename string
var flagBenchmark bool

func init() {
	flag.StringVar(&flagFilename, "f", "", "Path to word.list")
	flag.BoolVar(&flagBenchmark, "b", false, "Benchmark")
	flag.Parse()
}

func main() {
	if len(flagFilename) == 0 {
		fmt.Printf("Usage: %s -f filename\n", os.Args[0])
		os.Exit(1)
	}

	var start time.Time
	if flagBenchmark {
		start = time.Now()
	}

	s := lcwsolver.New()
	err := s.LoadWordList(flagFilename)
	if err != nil {
		log.Fatal(err.Error())
	}
	result := s.GetResult()

	if flagBenchmark {
		elapsed := time.Since(start)
		fmt.Println("Runtime:", elapsed.Nanoseconds()/1000000, "ms")
	}

	fmt.Println("Longest compound word:", result)
}
