package main

import (
	"Shyvana/modules/web/fingerprints"
	"fmt"
)

func main(){
	//web.LaunchWebScan()
	//body := "1111111111=.pHp111.phtm=lzzz.phP111"
	//re, _ := regexp.Compile(`(?i)=(.*?)111`)
	//res := re.FindAllStringSubmatch(body, -1)
	//fmt.Println(res)
	res := fingerprints.DetectCms()
	fmt.Println(res)
}