package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"math"
	"regexp"
)

func StrIsIn(item string, list []string)bool{
	for _, c := range(list){
		if c == item{
			return true
		}
	}
	return false
}

func StrLikelyIn(item string, list []string)bool{
	for _, c := range(list){
		match, _ := regexp.MatchString(`(?i)`+item, c)
		if match{
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

// Md5 func
func MD5(s string) (m string) {
	h := md5.New()
	_, _ = io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}