package main

import "testing"

func Average(arr []float64) float64 {
	var sum float64
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum / float64(len(arr))
}

func TestAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1,2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}