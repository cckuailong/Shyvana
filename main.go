package main

import (
	"Shyvana/utils"
	"Shyvana/vars"
	"fmt"
)

func main(){
	fmt.Println(vars.Webinfo.Web_url)
	body,status_code := utils.GetRespBody("https://www.douyu.com/hjj.html")
	fmt.Println(status_code)
	fmt.Println(utils.Is404(body))
	//web.LaunchWebScan()
	//body := `href="assets/css/bootstrap.min.css"`
	//re, _ := regexp.Compile(`(?i)href\s?=\s?["\|']#?(.*?)["\|']`)
	//uris := re.FindAllStringSubmatch(body, -1)
	//fmt.Println(uris)
}