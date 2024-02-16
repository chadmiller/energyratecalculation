package tariffs

import "time"


type gapoToureo16Calculator struct {
	costAccum CostCents
}

func (rat *gapoToureo16Calculator) String() string { return "GP TOU-REO-16" }

func (rat *gapoToureo16Calculator) Describe() string {
	return "Time-of-Use Residential Energy Only"
}

func (rat *gapoToureo16Calculator) Compute() CostCents { return (rat.costAccum + (46.03 * 30.5) + 2129/*OPT-5*/) * (1 + 0.030701/*MFF-10*/ + 0.015315/*DSM-R-13*/ + 0.041562/*NCCR-13*/) }

func (rat *gapoToureo16Calculator) addRow(when time.Time, usage UsagekWh, temperature TemperatureFahrenheit) CostCents {
	/*
	https://www.georgiapower.com/content/dam/georgia-power/pdfs/residential-pdfs/tariffs/2024/tou-reo-16.pdf
	*/
	var rate float32
	if when.Month() >= 6 && when.Month() <= 9 && when.Hour() >= 14 && when.Hour() <= 19 &&
		when.Weekday() >= time.Monday && when.Weekday() <= time.Friday {  // on-peak
		rate = 24.7091 + 4.3/*TOU-FCR-TP-4*/ + 6.6/*TOU-FCR-6*/
	} else {  // off peak
		rate = 6.6688 + 4.3/*TOU-FCR-TP-4*/ + 4.2/*TOU-FCR-6*/
	}

	rat.costAccum += CostCents(rate * float32(usage))
	return rat.Compute()
}
