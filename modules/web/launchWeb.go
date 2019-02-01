package web

import (
	"Shyvana/logger"
	"Shyvana/modules/web/fingerprints"
	"Shyvana/utils"
	"Shyvana/vars"
	"fmt"
	"io/ioutil"
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

func getRespBody()string{
	params := url.Values{}
	resp, err := utils.Http_req(vars.Webinfo.Web_url, params, "GET", vars.Headers)
	if err != nil{
		return ""
	}
	body,  _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func LaunchWebScan(){
	// Get the response headers with HEAD
	resp_header := getRespHeader()
	if resp_header == nil{
		logger.Log.Println("[Error][ HttpErr ] Get Response Headers Error")
	}
	// Get the response headers with OPTIONS
	resp_opt_header := getHttpMethod()
	if resp_opt_header == nil{
		logger.Log.Println("[Error][ HttpErr ] Get Http Method Error")
	}
	// Get the response body with GET
	resp_body := getRespBody()
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
	cs_lang := fingerprints.GetCsLang(resp_header, resp_body)
	fmt.Println(cs_lang)
}
