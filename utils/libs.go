package utils

import (
	"math"
)

func StrIsIn(item string, list []string)bool{
	for _, c := range(list){
		if c == item{
			return true
		}
	}
	return false
}

func MaxMap(map_l map[string]int)(string, int){
	res_k := ""
	res_v := math.MinInt64
	for k, v := range(map_l){
		if v > res_v{
			res_k = k
			res_v = v
		}
	}
	return res_k, res_v
}
