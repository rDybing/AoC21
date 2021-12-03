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
	hPos    int
	p1Depth int
	p2Depth int
	aim     int
	dir     directionT
	steps   int
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
			v.p1Depth = d[i-1].p1Depth
			v.p2Depth = d[i-1].p2Depth
			v.aim = d[i-1].aim
		}
		switch v.dir {
		case forward:
			v.hPos += v.steps
			v.p2Depth = v.p2Depth + (v.steps * v.aim)
		case up:
			v.p1Depth -= v.steps
			v.aim = v.aim - v.steps
		case down:
			v.p1Depth += v.steps
			v.aim = v.aim + v.steps
		}
		d[i] = v
	}
	last := d[len(d)-1]
	p1Sum := last.hPos * last.p1Depth
	p2Sum := last.hPos * last.p2Depth
	done := time.Now()
	diff := done.Sub(start)
	fmt.Printf("Part 1: hPos (%d) * p1Depth (%d) = %d\n", last.hPos, last.p1Depth, p1Sum)
	fmt.Printf("Part 2: hPos (%d) * p2Depth (%d) = %d\n", last.hPos, last.p2Depth, p2Sum)
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
