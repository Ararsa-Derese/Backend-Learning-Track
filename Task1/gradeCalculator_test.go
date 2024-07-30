package main 

import (
    "testing"
)

func TestCalculateAverage(t *testing.T) {
	grades := map[string]int{"Math": 90, "Science": 80, "English": 70}
	avg := calculateAverage(grades)
	if avg != 80 {
		t.Errorf("Expected 80, got %v", avg)
	}

}