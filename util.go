package smd

import (
	"strconv"
)

func ssto3f(v []string) (f [3]float64, err error) {
	var ff float64
	for i := 0; i < 3; i++ {
		ff, err = strconv.ParseFloat(v[i], 32)
		if err != nil {
			return
		}
		f[i] = float64(ff)
	}
	return
}

func ssto2f(v []string) (f [2]float64, err error) {
	var ff float64
	for i := 0; i < 2; i++ {
		ff, err = strconv.ParseFloat(v[i], 32)
		if err != nil {
			return
		}
		f[i] = float64(ff)
	}
	return
}
