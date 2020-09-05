package aqi

import "math"

const (
	no2GoodBreakpointLow           = 0.0
	no2GoodBreakpointHigh          = 53.0
	no2ModerateBreakpointLow       = 54.0
	no2ModerateBreakpointHigh      = 100.0
	no2SensitiveBreakpointLow      = 101.0
	no2SensitiveBreakpointHigh     = 360.0
	no2UnhealthyBreakpointLow      = 361.0
	no2UnhealthyBreakpointHigh     = 649.0
	no2VeryUnhealthyBreakpointLow  = 650.0
	no2VeryUnhealthyBreakpointHigh = 1249.0
	no2HazardousBreakpointLow      = 1250.0
	no2HazardousBreakpointHigh     = 1649.0
	no2VeryHazardousBreakpointLow  = 1650.0
	no2VeryHazardousBreakpointHigh = 2049.0
)

// NO2 contains concentration measurements for nitrogen dioxide in air in parts per billion.
type NO2 struct {
	Concentration float64
}

func (n NO2) findRangeAndCategory() (float64, float64, category) {
	c := math.Round(n.Concentration)
	if c >= no2GoodBreakpointLow && c <= no2GoodBreakpointHigh {
		return no2GoodBreakpointLow, no2GoodBreakpointHigh, categoryGood
	} else if c >= no2ModerateBreakpointLow && c <= no2ModerateBreakpointHigh {
		return no2ModerateBreakpointLow, no2ModerateBreakpointHigh, categoryModerate
	} else if c >= no2SensitiveBreakpointLow && c <= no2SensitiveBreakpointHigh {
		return no2SensitiveBreakpointLow, no2SensitiveBreakpointHigh, categorySensitive
	} else if c >= no2UnhealthyBreakpointLow && c <= no2UnhealthyBreakpointHigh {
		return no2UnhealthyBreakpointLow, no2UnhealthyBreakpointHigh, categoryUnhealthy
	} else if c >= no2VeryUnhealthyBreakpointLow && c <= no2VeryUnhealthyBreakpointHigh {
		return no2VeryUnhealthyBreakpointLow, no2VeryUnhealthyBreakpointHigh, categoryVeryUnhealthy
	} else if c >= no2HazardousBreakpointLow && c <= no2HazardousBreakpointHigh {
		return no2HazardousBreakpointLow, no2HazardousBreakpointHigh, categoryHazardous
	} else {
		return no2VeryHazardousBreakpointLow, no2VeryHazardousBreakpointHigh, categoryVeryHazardous
	}
}

func (n NO2) indexes() (float64, float64) {
	cLow, cHigh, _ := n.findRangeAndCategory()
	return cLow, cHigh
}

func (n NO2) category() category {
	_, _, category := n.findRangeAndCategory()
	return category
}

func (n NO2) value() float64 {
	return n.Concentration
}
