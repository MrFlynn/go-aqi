package aqi

import "math"

const (
	pm10GoodBreakpointLow           = 0.0
	pm10GoodBreakpointHigh          = 54.0
	pm10ModerateBreakpointLow       = 55.0
	pm10ModerateBreakpointHigh      = 154.0
	pm10SensitiveBreakpointLow      = 155.0
	pm10SensitiveBreakpointHigh     = 254.0
	pm10UnhealthyBreakpointLow      = 255.0
	pm10UnhealthyBreakpointHigh     = 354.0
	pm10VeryUnhealthyBreakpointLow  = 355.0
	pm10VeryUnhealthyBreakpointHigh = 424.0
	pm10HazardousBreakpointLow      = 425.0
	pm10HazardousBreakpointHigh     = 504.0
	pm10VeryHazardousBreakpointLow  = 505.0
	pm10VeryHazardousBreakpointHigh = 604.0
)

// PM10 contains concentration measurements for PM10 particulates in air in micrograms per meter cubed.
type PM10 struct {
	Concentration float64
}

func (p PM10) findRangeAndCategory() (float64, float64, category) {
	c := math.Round(p.Concentration)
	if c >= pm10GoodBreakpointLow && c <= pm10GoodBreakpointHigh {
		return pm10GoodBreakpointLow, pm10GoodBreakpointHigh, categoryGood
	} else if c >= pm10ModerateBreakpointLow && c <= pm10ModerateBreakpointHigh {
		return pm10ModerateBreakpointLow, pm10ModerateBreakpointHigh, categoryModerate
	} else if c >= pm10SensitiveBreakpointLow && c <= pm10SensitiveBreakpointHigh {
		return pm10SensitiveBreakpointLow, pm10SensitiveBreakpointHigh, categorySensitive
	} else if c >= pm10UnhealthyBreakpointLow && c <= pm10UnhealthyBreakpointHigh {
		return pm10UnhealthyBreakpointLow, pm10UnhealthyBreakpointHigh, categoryUnhealthy
	} else if c >= pm10VeryUnhealthyBreakpointLow && c <= pm10VeryUnhealthyBreakpointHigh {
		return pm10VeryUnhealthyBreakpointLow, pm10VeryUnhealthyBreakpointHigh, categoryVeryUnhealthy
	} else if c >= pm10HazardousBreakpointLow && c <= pm10HazardousBreakpointHigh {
		return pm10HazardousBreakpointLow, pm10HazardousBreakpointHigh, categoryHazardous
	} else {
		return pm10VeryHazardousBreakpointLow, pm10VeryHazardousBreakpointHigh, categoryVeryHazardous
	}
}

func (p PM10) indexes() (float64, float64) {
	cLow, cHigh, _ := p.findRangeAndCategory()
	return cLow, cHigh
}

func (p PM10) category() category {
	_, _, category := p.findRangeAndCategory()
	return category
}

func (p PM10) value() float64 {
	return p.Concentration
}
