package meanStdFunc

import "math"

func MeanFunc(freqs []float64) float64 {
	var sum float64
	for _, num := range freqs {
		sum += num
	}
	return float64(sum) / float64(len(freqs))

}

func STDDevFunc(nums []float64, mean float64) (float64, bool) {
	var sum float64
	for _, val := range nums {
		sum += (val - mean) * (val - mean)
	}
	variance := sum / float64(len(nums))
	c := math.Sqrt(variance)
	return c, true
}
