package math

func Average(arr []float64) float64 {
	var total float64 = 0
	for _, value := range arr {
		total += value
	}
	return total / float64(len(arr))
}
