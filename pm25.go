package aqi

import "github.com/shopspring/decimal"

const (
	pm25GoodBreakpointLow           = 0.0
	pm25GoodBreakpointHigh          = 12.0
	pm25ModerateBreakpointLow       = 12.1
	pm25ModerateBreakpointHigh      = 35.4
	pm25SensitiveBreakpointLow      = 35.5
	pm25SensitiveBreakpointHigh     = 55.4
	pm25UnhealthyBreakpointLow      = 55.5
	pm25UnhealthyBreakpointHigh     = 150.4
	pm25VeryUnhealthyBreakpointLow  = 150.5
	pm25VeryUnhealthyBreakpointHigh = 250.4
	pm25HazardousBreakpointLow      = 250.5
	pm25HazardousBreakpointHigh     = 350.4
	pm25VeryHazardousBreakpointLow  = 350.5
	pm25VeryHazardousBreakpointHigh = 500.4
)

// PM25 contains concentration measurements for PM2.5 particulates in air in micrograms per meter cubed.
type PM25 struct {
	Concentration float64
}

func (p PM25) findRangeAndCategory() (float64, float64, category) {
	c, _ := decimal.NewFromFloat(p.Concentration).Round(1).Float64()
	if c >= pm25GoodBreakpointLow && c <= pm25GoodBreakpointHigh {
		return pm25GoodBreakpointLow, pm25GoodBreakpointHigh, categoryGood
	} else if c >= pm25ModerateBreakpointLow && c <= pm25ModerateBreakpointHigh {
		return pm25ModerateBreakpointLow, pm25ModerateBreakpointHigh, categoryModerate
	} else if c >= pm25SensitiveBreakpointLow && c <= pm25SensitiveBreakpointHigh {
		return pm25SensitiveBreakpointLow, pm25SensitiveBreakpointHigh, categorySensitive
	} else if c >= pm25UnhealthyBreakpointLow && c <= pm25UnhealthyBreakpointHigh {
		return pm25UnhealthyBreakpointLow, pm25UnhealthyBreakpointHigh, categoryUnhealthy
	} else if c >= pm25VeryUnhealthyBreakpointLow && c <= pm25VeryUnhealthyBreakpointHigh {
		return pm25VeryUnhealthyBreakpointLow, pm25VeryUnhealthyBreakpointHigh, categoryVeryUnhealthy
	} else if c >= pm25HazardousBreakpointLow && c <= pm25HazardousBreakpointHigh {
		return pm25HazardousBreakpointLow, pm25HazardousBreakpointHigh, categoryHazardous
	} else {
		return pm25VeryHazardousBreakpointLow, pm25VeryHazardousBreakpointHigh, categoryVeryHazardous
	}
}

func (p PM25) indexes() (float64, float64) {
	cLow, cHigh, _ := p.findRangeAndCategory()
	return cLow, cHigh
}

func (p PM25) category() category {
	_, _, category := p.findRangeAndCategory()
	return category
}

func (p PM25) value() float64 {
	return p.Concentration
}
