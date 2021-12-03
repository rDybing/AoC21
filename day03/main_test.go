package main

import "testing"

func TestAddBitSum(t *testing.T) {
	d := []dataT{
		{bin: 0b100100110110},
		{bin: 0b101110110110},
		{bin: 0b010100010100},
		{bin: 0b011001110000},
		{bin: 0b000000000111},
		{bin: 0b000010110001},
		{bin: 0b001111000001},
		{bin: 0b100010000001},
		{bin: 0b010100110011},
		{bin: 0b010000010110},
		{bin: 0b010000000011},
		{bin: 0b010101001000},
		{bin: 0b011011101100},
	}
	r := rateT{
		one:     [12]int{3, 7, 4, 6, 5, 4, 6, 7, 2, 6, 6, 6},
		zero:    [12]int{10, 6, 9, 7, 8, 9, 7, 6, 11, 7, 7, 7},
		gamma:   0b0000010000010000,
		epsilon: 0b0000101111101111,
	}
	tt := struct {
		name  string
		data  []dataT
		rateA rateT
		sumA  uint32
	}{
		name:  "374654672666",
		data:  d,
		rateA: r,
		sumA:  3177200,
	}
	t.Run(tt.name, func(t *testing.T) {
		var rateQ rateT
		for _, v := range tt.data {
			rateQ.addBitSum(v)
		}
		for i, v := range rateQ.zero {
			if v == 12 {
				rateQ.zero[i] = 0
			}
		}
		rateQ.getGammaEpsilon()
		sumQ := uint32(rateQ.gamma) * uint32(rateQ.epsilon)
		if rateQ != tt.rateA {
			t.Fatalf("\n%s Expected: %d\n%+v\ngot: %d\n%+v\n", tt.name, tt.sumA, tt.rateA, sumQ, rateQ)
		}
	})
}
