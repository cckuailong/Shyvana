package main

import (
	"Shyvana/modules/web/crawler"
	"Shyvana/modules/web/fingerprints"
)

func main(){
	//web.LaunchWebScan()
	//body := "1111111111=.pHp111.phtm=lzzz.phP111"
	//re, _ := regexp.Compile(`(?i)=(.*?)111`)
	//res := re.FindAllStringSubmatch(body, -1)
	//fmt.Println(res)
	ip_l := fingerprints.GetIPs()
	crawler.Crawl(ip_l)
}