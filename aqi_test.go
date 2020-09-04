package aqi

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var floatComparer = cmp.Comparer(func(x, y float64) bool {
	delta := math.Abs(x - y)
	mean := math.Abs(x+y) / 2.0

	return delta/mean < 0.00001
})

func TestPM25(t *testing.T) {
	result, err := Calculate(PM25{26.4})
	if err != nil {
		t.Errorf("got unexpected error: %s", err)
	}

	if !cmp.Equal(result.AQI, 81.073, floatComparer) {
		t.Errorf("expected AQI of 81.073, got %.3f", result.AQI)
	}
}

func TestPM25Extreme(t *testing.T) {
	result, err := Calculate(PM25{500.0})
	if err != nil {
		t.Errorf("got unexpected error: %s", err)
	}

	if !cmp.Equal(result.AQI, 499.681, floatComparer) {
		t.Errorf("expected AQI for 499.736, got %.3f", result.AQI)
	}
}

func TestMultipleInputs(t *testing.T) {
	result, err := Calculate(PM25{30.0}, PM25{10.0})
	if err != nil {
		t.Errorf("got unexpected error: %s", err)
	}

	if !cmp.Equal(result.AQI, 88.643, floatComparer) {
		t.Errorf("expected AQI of 88.643, got %.3f", result.AQI)
	}
}

func TestErrorCondition(t *testing.T) {
	_, err := Calculate(PM25{-10.0})
	if err == nil {
		t.Error("expected error, got nil")
	}
}