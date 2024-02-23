package lc2469

func convertTemperature(celsius float64) []float64 {
	k := celsius + 273.15
	f := celsius*1.8 + 32
	return []float64{k, f}
}
