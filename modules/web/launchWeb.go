package web

import (
	"Shyvana/logger"
	"Shyvana/utils"
	"Shyvana/vars"
)

func LaunchWebScan(){
	// Get the response headers with HEAD
	//resp_header := utils.GetRespHeader()
	//if resp_header == nil{
	//	logger.Log.Println("[Error][ HttpErr ] Get Response Headers Error")
	//}
	//fmt.Println(resp_header)

	// Get the response headers with OPTIONS
	//resp_opt_header := utils.GetHttpMethod()
	//if resp_opt_header == nil{
	//	logger.Log.Println("[Error][ HttpErr ] Get Http Method Error")
	//}

	//Get the response body with GET
	resp_body := utils.GetRespBody(vars.Webinfo.Web_url)
	if len(resp_body) == 0{
		logger.Log.Println("[ Warinng ][ HttpWarn ] Get Http Body Error or Empty Body")
	}

	// Get the Web Server, Like Apache, Nginx and so on
	// Empty: ""
	//serv_info := fingerprints.GetWebServ(resp_header)
	//fmt.Println(serv_info)

	// Get the Http Options
	// Empty: len() == 0
	//http_method := fingerprints.GetHttpMethod(resp_opt_header)
	//fmt.Println(http_method)

	// Verify the lang (php and so on)
	//cs_lang := fingerprints.GetCsLang(resp_header, resp_body)
	//fmt.Println(cs_lang)

	// Detect Cookie Security
	//cookieres := fingerprints.DetectCookieSec(resp_header["Set-Cookie"][0])
	//fmt.Println(cookieres)

	// Detect Headers Security
	//headerres := fingerprints.DetectHeaderSec(resp_header)
	//fmt.Println(headerres)

	// Detect Waf
	//res := fingerprints.DetectWaf(resp_header, resp_body)
	//fmt.Println(res)

	// Detect cdn
	//res := fingerprints.DetectCdn(resp_header)
	//fmt.Println(res)

	// Detect frontend
	//res, err := fingerprints.DetectFrontEnd(resp_body)
	//fmt.Println(res, err)

	// Get Whois Info
	//fingerprints.GetWhoisInfo()

}
