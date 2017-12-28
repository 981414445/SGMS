package util

import (
	"fmt"
	"testing"
)

func TestStat(t *testing.T) {
	fmt.Printf("%#v\n", Stat.GetStats([]float64{1, 0, 1, 2, 1, 2, 3, 2, 3}))
}
