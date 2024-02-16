package tariffs

import "time"


type gapoR28Calculator struct {
	runningUsage UsagekWh
	costAccum CostCents
}

func (rat *gapoR28Calculator) String() string { return "GP R-28" }

func (rat *gapoR28Calculator) Describe() string {
	return "Residential"
}

func (rat *gapoR28Calculator) Compute() CostCents { return (rat.costAccum + (46.03 * 30.5) + 2129/*OPT-5*/) * (1 + 0.030701/*MFF-10*/ + 0.015315/*DSM-R-13*/ + 0.041562/*NCCR-13*/) }

func (rat *gapoR28Calculator) addRow(when time.Time, usage UsagekWh, temperature TemperatureFahrenheit) CostCents {
	/*
	https://www.georgiapower.com/content/dam/georgia-power/pdfs/residential-pdfs/tariffs/2024/r-28.pdf
	*/
	rat.runningUsage += usage

	var rate float32
	if when.Month() >= 6 && when.Month() <= 9 {  // summer
		if rat.runningUsage <= 650 {
			rate = 7.1255 + 4.5/*FCR-26*/
		} else if rat.runningUsage <= 650+350 {
			rate = 11.8349 + 4.5/*FCR-26*/
		} else {
			rate = 12.2493 + 4.5/*FCR-26*/
		}
	} else {  // not summer
		rate = 6.6688 + 4.3/*FCR-26*/
	}

	rat.costAccum += CostCents(rate * float32(usage))
	return rat.Compute()
}
