package aqi

import "math"

const (
	so2GoodBreakpointLow           = 0.0
	so2GoodBreakpointHigh          = 35.0
	so2ModerateBreakpointLow       = 36.0
	so2ModerateBreakpointHigh      = 75.0
	so2SensitiveBreakpointLow      = 76.0
	so2SensitiveBreakpointHigh     = 185.0
	so2UnhealthyBreakpointLow      = 186.0
	so2UnhealthyBreakpointHigh     = 304.0
	so2VeryUnhealthyBreakpointLow  = 305.0
	so2VeryUnhealthyBreakpointHigh = 604.0
	so2HazardousBreakpointLow      = 605.0
	so2HazardousBreakpointHigh     = 804.0
	so2VeryHazardousBreakpointLow  = 805.0
	so2VeryHazardousBreakpointHigh = 1004.0
)

// SO2 ontains concentration measurements for sulfur dioxide measurements in air in parts per billion.
type SO2 struct {
	Concentration float64
}

func (s SO2) findRangeAndCategory() (float64, float64, category) {
	c := math.Round(s.Concentration)
	if c >= so2GoodBreakpointLow && c <= so2GoodBreakpointHigh {
		return so2GoodBreakpointLow, so2GoodBreakpointHigh, categoryGood
	} else if c >= so2ModerateBreakpointLow && c <= so2ModerateBreakpointHigh {
		return so2ModerateBreakpointLow, so2ModerateBreakpointHigh, categoryModerate
	} else if c >= so2SensitiveBreakpointLow && c <= so2SensitiveBreakpointHigh {
		return so2SensitiveBreakpointLow, so2SensitiveBreakpointHigh, categorySensitive
	} else if c >= so2UnhealthyBreakpointLow && c <= so2UnhealthyBreakpointHigh {
		return so2UnhealthyBreakpointLow, so2UnhealthyBreakpointHigh, categoryUnhealthy
	} else if c >= so2VeryUnhealthyBreakpointLow && c <= so2VeryUnhealthyBreakpointHigh {
		return so2VeryUnhealthyBreakpointLow, so2VeryUnhealthyBreakpointHigh, categoryVeryUnhealthy
	} else if c >= so2HazardousBreakpointLow && c <= so2HazardousBreakpointHigh {
		return so2HazardousBreakpointLow, so2HazardousBreakpointHigh, categoryHazardous
	} else {
		return so2VeryHazardousBreakpointLow, so2VeryHazardousBreakpointHigh, categoryVeryHazardous
	}
}

func (s SO2) indexes() (float64, float64) {
	cLow, cHigh, _ := s.findRangeAndCategory()
	return cLow, cHigh
}

func (s SO2) category() category {
	_, _, category := s.findRangeAndCategory()
	return category
}

func (s SO2) value() float64 {
	return s.Concentration
}
