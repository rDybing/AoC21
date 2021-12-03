package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

type directionT int

const (
	forward directionT = iota
	up
	down
)

type dataT struct {
	hPos  int
	depth int
	dir   directionT
	steps int
}

func main() {
	d, err := importData()
	if err != nil {
		log.Panicf("Failed loading data file: %v\n", err)
	}
	start := time.Now()
	for i, v := range d {
		if i > 0 {
			v.hPos = d[i-1].hPos
			v.depth = d[i-1].depth
		}
		switch v.dir {
		case forward:
			v.hPos += v.steps
		case up:
			v.depth -= v.steps
		case down:
			v.depth += v.steps
		}
		d[i] = v
	}
	last := d[len(d)-1]
	p1Sum := last.hPos * last.depth
	done := time.Now()
	diff := done.Sub(start)
	fmt.Printf("Part 1: hPos (%d) * depth (%d) = %d\n", last.hPos, last.depth, p1Sum)
	//fmt.Printf("Part 2: Found %d three-sum sliding increases\n", sumCount)
	fmt.Printf("Execution time: %f\n", diff.Seconds())
}

func importData() ([]dataT, error) {
	var d []dataT
	f, err := ioutil.ReadFile("./data.txt")
	if err != nil {
		return d, err
	}
	arr := strings.Split(string(f), "\n")
	for _, v := range arr {
		content := strings.Split(v, " ")
		direction := enumDirection(content[0])
		steps, _ := strconv.Atoi(content[1])
		dT := dataT{
			dir:   direction,
			steps: steps,
		}
		d = append(d, dT)
	}
	fmt.Printf("loaded %d data-points\n", len(d))
	return d, nil
}

func enumDirection(in string) directionT {
	var out directionT
	switch in {
	case "forward":
		out = forward
	case "up":
		out = up
	case "down":
		out = down
	}
	return out
}
