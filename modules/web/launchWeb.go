package web

import (
	"Shyvana/logger"
	"Shyvana/modules/web/fingerprints"
	"Shyvana/utils"
	"Shyvana/vars"
	"fmt"
	"net/http"
	"net/url"
)

func getRespHeader()http.Header{
	params := url.Values{}
	resp, err := utils.Http_req(vars.Webinfo.Web_url, params, "HEAD", vars.Headers)
	if err != nil{
		return nil
	}
	return resp.Header
}

func getHttpMethod()http.Header{
	params := url.Values{}
	resp, err := utils.Http_req(vars.Webinfo.Web_url, params, "OPTIONS", vars.Headers)
	if err != nil{
		return nil
	}
	return resp.Header
}

func LaunchWebScan(){
	// Get the response headers
	resp_header := getRespHeader()
	if resp_header == nil{
		logger.Log.Println("[Error][ HttpErr ] Get Response Headers Error")
	}
	// Get the response headers with options
	resp_opt_header := getHttpMethod()
	if resp_opt_header == nil{
		logger.Log.Println("[Error][ HttpErr ] Get Http Method Error")
	}

	// Get the Web Server, Like Apache, Nginx and so on
	// Empty: ""
	serv_info := fingerprints.GetWebServ(resp_header)
	fmt.Println(serv_info)
	// Get the Http Options
	// Empty: len() == 0
	http_method := fingerprints.GetHttpMethod(resp_opt_header)
	fmt.Println(http_method)
}
