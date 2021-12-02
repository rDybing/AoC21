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
	depth    int
	increase bool
}

func main() {
	d, err := importData()
	if err != nil {
		log.Panicf("Failed loading data file: %v\n", err)
	}
	var incCount int
	start := time.Now()
	for i, v := range d {
		if i != 0 {
			if v.depth > d[i-1].depth {
				v.increase = true
				incCount++
			}
		} else {
			v.increase = false
		}
		d[i] = v
	}
	done := time.Now()
	diff := done.Sub(start)
	fmt.Printf("Found %d depth increases in %f seconds\n", incCount, diff.Seconds())
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
			depth:    depth,
			increase: false,
		}
		d = append(d, dT)
	}
	fmt.Printf("loaded %d data-points\n", len(d))
	return d, nil
}
