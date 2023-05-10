package constant

import "math"

const AppName = "AHP TPS Location Services "

const DbDefaultCreateBy = "system"

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
