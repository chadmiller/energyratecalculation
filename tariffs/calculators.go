package tariffs

import "time"


type CostCents float32

type UsagekWh float32

type TemperatureFahrenheit float32

type Calculator interface {
	addRow(when time.Time, usage UsagekWh, temperature TemperatureFahrenheit) CostCents
	Compute() CostCents
	Describe() string
	String() string
	//IdentifyForFiltering() string
}


func GetCalculators(matches string) []Calculator {
	ret := make([]Calculator, 0, 10)

	ret = append(ret, &(gapoR28Calculator{}))
	ret = append(ret, &gapoToupev12Calculator{})
	ret = append(ret, &gapoToureo16Calculator{})

	return ret
}

func AddRowToAll(calcs []Calculator, t time.Time, usage UsagekWh, temp TemperatureFahrenheit) {
	for _, calc := range(calcs) {
		calc.addRow(t, usage, temp)
	}
}
