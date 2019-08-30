package main

import (
	"Shyvana/modules/web/vul"
	"fmt"
)

func main(){
	fmt.Println(vul.Check_ssrf("https://baidu.com/"))
	//web.LaunchWebScan()
	//body := `href="assets/css/bootstrap.min.css"`
	//re, _ := regexp.Compile(`(?i)href\s?=\s?["\|']#?(.*?)["\|']`)
	//uris := re.FindAllStringSubmatch(body, -1)
	//fmt.Println(uris)
}