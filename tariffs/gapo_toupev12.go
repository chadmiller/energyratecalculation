package tariffs

import "time"


type gapoToupev12Calculator struct {
	costAccum CostCents
}

func (rat *gapoToupev12Calculator) String() string { return "GP TOU-PEV-12" }

func (rat *gapoToupev12Calculator) Describe() string {
	return "Time-of-Use Plug-in Electric Vehicle"
}

func (rat *gapoToupev12Calculator) Compute() CostCents { return (rat.costAccum + (46.03 * 30.5) + 2129/*OPT-5*/) * (1 + 0.030701/*MFF-10*/ + 0.015315/*DSM-R-13*/ + 0.041562/*NCCR-13*/) }

func (rat *gapoToupev12Calculator) addRow(when time.Time, usage UsagekWh, temperature TemperatureFahrenheit) CostCents {
	/*
	https://www.georgiapower.com/residential/billing-and-rate-plans/pricing-and-rate-plans/plug-in-ev.html
	https://www.georgiapower.com/content/dam/georgia-power/pdfs/residential-pdfs/tariffs/2024/tou-pev-12.pdf
	*/
	var rate float32
	if when.Hour() >= 23 || when.Hour() < 7 {  // in super off-peak?
		rate = 1.8160 + 4.3 + 3.7/*TOU-FCR-TP-4*/
	} else if when.Month() >= 6 && when.Month() <= 9 && when.Hour() >= 14 && when.Hour() <= 19 {  // on-peak
		rate = 24.7091 + 4.3 + 6.5/*TOU-FCR-TP-4*/ + 6.6/*TOU-FCR-6*/
	} else {  
		rate = 8.4469 + 4.3 + 4.3/*TOU-FCR-TP-4*/ + 4.2/*TOU-FCR-6*/
	}

	rat.costAccum += CostCents(rate * float32(usage))
	return rat.Compute()
}

