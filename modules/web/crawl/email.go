package crawl

import (
	"Shyvana/utils"
	"regexp"
)

func FetchEmail(body string, fetched_l []string)[]string{
	re, _ := regexp.Compile(`[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+`)
	res_l := re.FindAllString(body, -1)
	for _, c := range(res_l){
		if !utils.StrIsIn(c, fetched_l){
			fetched_l = append(fetched_l, c)
		}
	}
	return fetched_l
}
