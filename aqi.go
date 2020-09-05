package aqi

import (
	"fmt"
	"image/color"
)

const (
	goodBreakpointLow           = 0
	goodBreakpointHigh          = 50
	moderateBreakpointLow       = 51
	moderateBreakpointHigh      = 100
	sensitiveBreakpointLow      = 101
	sensitiveBreakpointHigh     = 150
	unhealthyBreakpointLow      = 151
	unhealthyBreakpointHigh     = 200
	veryUnhealthyBreakpointLow  = 201
	veryUnhealthyBreakpointHigh = 300
	hazardousBreakpointLow      = 301
	hazardousBreakpointHigh     = 400
	veryHazardousBreakpointLow  = 401
	veryHazardousBreakpointHigh = 500
)

type category int8

const (
	categoryGood category = iota
	categoryModerate
	categorySensitive
	categoryUnhealthy
	categoryVeryUnhealthy
	categoryHazardous
	categoryVeryHazardous
)

// Index represents the different levels of AQI (i.e good, moderate, etc.) with associated
// metadata.
type Index struct {
	Name  string
	Color color.RGBA
	Low   int
	High  int
}

// Default indices for all EPA AQI levels, each containing their level's respective metadata.
var (
	Good = Index{
		Name:  "Good",
		Color: color.RGBA{0, 228, 0, 255},
		Low:   goodBreakpointLow,
		High:  goodBreakpointHigh,
	}
	Moderate = Index{
		Name:  "Moderate",
		Color: color.RGBA{255, 255, 0, 255},
		Low:   moderateBreakpointLow,
		High:  moderateBreakpointHigh,
	}
	Sensitive = Index{
		Name:  "Unhealthy for Sensitive Groups",
		Color: color.RGBA{255, 126, 0, 255},
		Low:   sensitiveBreakpointLow,
		High:  sensitiveBreakpointHigh,
	}
	Unhealthy = Index{
		Name:  "Unhealthy",
		Color: color.RGBA{255, 0, 0, 255},
		Low:   unhealthyBreakpointLow,
		High:  unhealthyBreakpointHigh,
	}
	VeryUnhealthy = Index{
		Name:  "Very Unhealthy",
		Color: color.RGBA{153, 0, 76, 255},
		Low:   veryUnhealthyBreakpointLow,
		High:  veryUnhealthyBreakpointHigh,
	}
	Hazardous = Index{
		Name:  "Hazardous",
		Color: color.RGBA{125, 0, 35, 255},
		Low:   hazardousBreakpointLow,
		High:  hazardousBreakpointHigh,
	}
	VeryHazardous = Index{
		Name:  "Hazardous",
		Color: color.RGBA{125, 0, 35, 255},
		Low:   veryHazardousBreakpointLow,
		High:  veryHazardousBreakpointHigh,
	}
)

// Measurement type are the functions require to calculate the AQI value.
type Measurement interface {
	indexes() (float64, float64)
	category() category
	value() float64
}

// Result contains the AQI one of the predefined Index values.
type Result struct {
	AQI   float64
	Index Index
}

func indexFromCategory(c category) (Index, error) {
	switch c {
	case categoryGood:
		return Good, nil
	case categoryModerate:
		return Moderate, nil
	case categorySensitive:
		return Sensitive, nil
	case categoryUnhealthy:
		return Unhealthy, nil
	case categoryVeryUnhealthy:
		return VeryUnhealthy, nil
	case categoryHazardous:
		return Hazardous, nil
	case categoryVeryHazardous:
		return VeryHazardous, nil
	default:
		return Index{}, fmt.Errorf("could not find index for category %v", c)
	}
}

// Calculate determines the AQI from the given measurements. The largest AQI value is always returned.
func Calculate(ms ...Measurement) (Result, error) {
	var value float64
	var index Index

	for _, m := range ms {
		if m.value() < 0.0 {
			return Result{}, fmt.Errorf("measurement for %T cannot be less than 0", m)
		}

		tmpIndex, err := indexFromCategory(m.category())
		if err != nil {
			return Result{}, err
		}

		cLow, cHigh := m.indexes()
		tmpValue := ((float64(tmpIndex.High)-float64(tmpIndex.Low))/(cHigh-cLow))*(m.value()-cLow) + float64(tmpIndex.Low)

		if tmpValue > value {
			value = tmpValue
			index = tmpIndex
		}
	}

	return Result{
		AQI:   value,
		Index: index,
	}, nil
}
