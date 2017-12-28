package util

import (
	"math"
	"sort"
)

//statistic methods
var Stat = _Stat{}

type StatSummary struct {
	Numbers []float64
	Mean    float64
	Median  float64
	Modes   []float64
	StdDev  float64
}

type _Stat struct {
}

func (m _Stat) GetStats(numbers []float64) (stats StatSummary) {
	stats.Numbers = numbers
	sort.Float64s(stats.Numbers)
	stats.Mean = m.Sum(numbers) / float64(len(numbers))
	stats.Median = m.Median(numbers)
	stats.Modes = m.Mode(numbers)
	stats.StdDev = m.StdDev(numbers, stats.Mean)
	return stats
}
func (m _Stat) Sum(numbers []float64) (total float64) {
	for _, x := range numbers {
		total += x
	}
	return total
}
func (m _Stat) Mean(numbers []float64) float64 {
	return m.Sum(numbers) / float64(len(numbers))
}
func (m _Stat) Median(numbers []float64) float64 {
	middle := len(numbers) / 2
	result := numbers[middle]
	if len(numbers)%2 == 0 {
		result = (result + numbers[middle-1]) / 2
	}
	return result
}
func (m _Stat) Mode(numbers []float64) (modes []float64) {
	frequencies := make(map[float64]int, len(numbers))
	highestFrequency := 0
	for _, x := range numbers {
		frequencies[x]++
		if frequencies[x] > highestFrequency {
			highestFrequency = frequencies[x]
		}
	}
	for x, frequency := range frequencies {
		if frequency == highestFrequency {
			modes = append(modes, x)
		}
	}
	if highestFrequency == 1 || len(modes) == len(numbers) {
		modes = modes[:0] // Or: modes = []float64{}
	}
	sort.Float64s(modes)
	return modes
}

func (m _Stat) StdDev(numbers []float64, mean float64) float64 {
	total := 0.0
	for _, number := range numbers {
		total += math.Pow(number-mean, 2)
	}
	variance := total / float64(len(numbers)-1)
	return math.Sqrt(variance)
}
