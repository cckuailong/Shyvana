package utils

import (
	"Shyvana/logger"
	"Shyvana/vars"
	"crypto/tls"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// http request func
func Http_req(uri string, value url.Values, method string, headers map[string]string, redirect bool) (*http.Response, error){
	var pv io.Reader
	var client *http.Client
	if value != nil{
		post_value := value.Encode()
		pv = strings.NewReader(post_value)
	}else{
		pv=nil
	}
	req, err := http.NewRequest(method, uri, pv)
	if err != nil{
		return nil, err
	}
	if headers != nil{
		for k,v := range(headers){
			req.Header.Add(k, v)
		}
	}
	// redirect option
	if redirect{
		client = &http.Client{}
	}else{
		client = &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
	}
	// ignore https verify
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	resp, err := client.Do(req)
	if err != nil{
		return nil, err
	}
	return resp, nil
}

func GetRespHeader()http.Header{
	params := url.Values{}
	resp, err := Http_req(vars.Webinfo.Web_url, params, "HEAD", vars.Headers, true)
	if err != nil{
		return nil
	}
	return resp.Header
}

func GetRespHeaderNoRedirect(uri string)(http.Header,int){
	params := url.Values{}
	resp, err := Http_req(uri, params, "HEAD", vars.Headers, false)
	if err != nil{
		return nil, 0
	}
	return resp.Header, resp.StatusCode
}

func GetHttpMethod()http.Header{
	params := url.Values{}
	resp, err := Http_req(vars.Webinfo.Web_url, params, "OPTIONS", vars.Headers, true)
	if err != nil{
		return nil
	}
	return resp.Header
}

func GetRespBody(uri string) (string, int){
	params := url.Values{}
	resp, err := Http_req(uri, params, "GET", vars.Headers,true)
	defer resp.Body.Close()
	if err != nil{
		logger.Log.Println("%v", err)
		return "", -1
	}
	body,  _ := ioutil.ReadAll(resp.Body)

	return string(body), resp.StatusCode
}
