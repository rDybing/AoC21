package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

type dataT struct {
	bin uint16
}

type rateT struct {
	one     [12]int
	zero    [12]int
	gamma   uint16
	epsilon uint16
}

func main() {
	d, err := importData()
	if err != nil {
		log.Panicf("Failed loading data file: %v\n", err)
	}
	var rate rateT
	start := time.Now()
	for _, v := range d {
		rate.addBitSum(v)
	}
	rate.getGammaEpsilon()
	sum := uint32(rate.gamma) * uint32(rate.epsilon)
	done := time.Now()
	diff := done.Sub(start)
	fmt.Printf("Part 1: Gamma (%012b) %d * Epsilon(%012b) %d = (%016b) %d\n%+v\n",
		rate.gamma, rate.gamma, rate.epsilon, rate.epsilon, sum, sum, rate)
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
		dT := dataT{
			bin: stringToByte(v),
		}
		d = append(d, dT)
	}
	fmt.Printf("loaded %d data-points\n", len(d))
	return d, nil
}

func (r *rateT) addBitSum(in dataT) {
	for i := 11; i >= 0; i-- {
		if in.bin&1 == 1 {
			r.one[i]++
		} else {
			r.zero[i]++
		}
		in.bin = in.bin >> 1
	}
}

func (r *rateT) getGammaEpsilon() {
	for i := 0; i < 12; i++ {
		if r.one[i] > r.zero[i] {
			r.gamma = r.gamma + 1
		} else {
			r.epsilon = r.epsilon + 1
		}
		if i != 11 {
			r.gamma = r.gamma << 1
			r.epsilon = r.epsilon << 1
		}
	}
}

func stringToByte(in string) uint16 {
	var out uint16
	for _, v := range in {
		out = out << 1
		if v == '1' {
			out++
		}
	}
	return out
}
