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

	if !cmp.Equal(result.Index, Moderate) {
		t.Errorf("expected Moderate{}, got %v", result.Index)
	}
}

func TestPM10(t *testing.T) {
	result, err := Calculate(PM10{160})
	if err != nil {
		t.Errorf("got unexpected error: %s", err)
	}

	if !cmp.Equal(result.AQI, 103.474, floatComparer) {
		t.Errorf("expected AQI of 103.474, got %.3f", result.AQI)
	}

	if !cmp.Equal(result.Index, Sensitive) {
		t.Errorf("expected Sensitive{}, got %v", result.Index)
	}
}

func TestCO(t *testing.T) {
	result, err := Calculate(CO{31})
	if err != nil {
		t.Errorf("got unexpected error: %s", err)
	}

	if !cmp.Equal(result.AQI, 306.0, floatComparer) {
		t.Errorf("expected AQI of 306.0, got %.3f", result.AQI)
	}

	if !cmp.Equal(result.Index, Hazardous) {
		t.Errorf("expected Hazardous{}, got %v", result.Index)
	}
}

func TestSO2(t *testing.T) {
	result, err := Calculate(SO2{48.0})
	if err != nil {
		t.Errorf("got unexpected error: %s", err)
	}

	if !cmp.Equal(result.AQI, 66.077, floatComparer) {
		t.Errorf("expected AQI of 66.077, got %.3f", result.AQI)
	}

	if !cmp.Equal(result.Index, Moderate) {
		t.Errorf("expected Moderate{}, got %v", result.Index)
	}
}

func TestNO2(t *testing.T) {
	result, err := Calculate(NO2{20.0})
	if err != nil {
		t.Errorf("got unexpected error: %s", err)
	}

	if !cmp.Equal(result.AQI, 18.868, floatComparer) {
		t.Errorf("expected AQI of 18.868, got %.3f", result.AQI)
	}

	if !cmp.Equal(result.Index, Good) {
		t.Errorf("expected Good{}, got %v", result.Index)
	}
}

func TestPM25Extreme(t *testing.T) {
	result, err := Calculate(PM25{500.0})
	if err != nil {
		t.Errorf("got unexpected error: %s", err)
	}

	if !cmp.Equal(result.AQI, 499.736, floatComparer) {
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