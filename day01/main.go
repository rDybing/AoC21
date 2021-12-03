package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

type dataT struct {
	depth int
	sum   int
}

func main() {
	d, err := importData()
	if err != nil {
		log.Panicf("Failed loading data file: %v\n", err)
	}
	var incCount, sumCount int
	start := time.Now()
	for i, v := range d {
		if i != 0 {
			if v.depth > d[i-1].depth {
				incCount++
			}
		}
		if i > 1 {
			v.sum = v.depth + d[i-1].depth + d[i-2].depth
		}
		if i > 2 {
			if v.sum > d[i-1].sum {
				sumCount++
			}
		}
		d[i] = v
	}
	done := time.Now()
	diff := done.Sub(start)
	fmt.Printf("Part 1: Found %d depth increases\n", incCount)
	fmt.Printf("Part 2: Found %d three-sum sliding increases\n", sumCount)
	fmt.Printf("Execution time: %d µSeconds\n", diff.Microseconds())
}

func importData() ([]dataT, error) {
	var d []dataT
	f, err := ioutil.ReadFile("./data.txt")
	if err != nil {
		return d, err
	}
	arr := strings.Split(string(f), "\n")
	for _, v := range arr {
		depth, _ := strconv.Atoi(v)
		dT := dataT{
			depth: depth,
			sum:   0,
		}
		d = append(d, dT)
	}
	fmt.Printf("loaded %d data-points\n", len(d))
	return d, nil
}
