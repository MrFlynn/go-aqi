package aqi

const (
	coGoodBreakpointLow           = 0.0
	coGoodBreakpointHigh          = 4.4
	coModerateBreakpointLow       = 4.5
	coModerateBreakpointHigh      = 9.4
	coSensitiveBreakpointLow      = 9.5
	coSensitiveBreakpointHigh     = 12.4
	coUnhealthyBreakpointLow      = 12.5
	coUnhealthyBreakpointHigh     = 15.4
	coVeryUnhealthyBreakpointLow  = 15.5
	coVeryUnhealthyBreakpointHigh = 30.4
	coHazardousBreakpointLow      = 30.5
	coHazardousBreakpointHigh     = 40.4
	coVeryHazardousBreakpointLow  = 40.5
	coVeryHazardousBreakpointHigh = 50.4
)

type CO struct {
	Concentration float64
}

func (c CO) findRangeAndCategory() (float64, float64, category) {
	v := c.Concentration
	if v >= coGoodBreakpointLow && v <= coGoodBreakpointHigh {
		return coGoodBreakpointLow, coGoodBreakpointHigh, categoryGood
	} else if v >= coModerateBreakpointLow && v <= coModerateBreakpointHigh {
		return coModerateBreakpointLow, coModerateBreakpointHigh, categoryModerate
	} else if v >= coSensitiveBreakpointLow && v <= coSensitiveBreakpointHigh {
		return coSensitiveBreakpointLow, coSensitiveBreakpointHigh, categorySensitive
	} else if v >= coUnhealthyBreakpointLow && v <= coUnhealthyBreakpointHigh {
		return coUnhealthyBreakpointLow, coUnhealthyBreakpointHigh, categoryUnhealthy
	} else if v >= coVeryUnhealthyBreakpointLow && v <= coVeryUnhealthyBreakpointHigh {
		return coVeryUnhealthyBreakpointLow, coVeryUnhealthyBreakpointHigh, categoryVeryUnhealthy
	} else if v >= coHazardousBreakpointLow && v <= coHazardousBreakpointHigh {
		return coHazardousBreakpointLow, coHazardousBreakpointHigh, categoryHazardous
	} else {
		return coVeryHazardousBreakpointLow, coVeryHazardousBreakpointHigh, categoryVeryHazardous
	}
}

func (c CO) Range() (float64, float64) {
	cLow, cHigh, _ := c.findRangeAndCategory()
	return cLow, cHigh
}

func (c CO) Category() category {
	_, _, category := c.findRangeAndCategory()
	return category
}

func (c CO) Value() float64 {
	return c.Concentration
}
