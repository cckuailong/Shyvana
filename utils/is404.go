package utils

import (
	"Shyvana/vars"
	"github.com/cckuailong/simHtml/simHtml"
)

func Get404() string{
	// 404 Page
	uri_404 := vars.Webinfo.Web_url+"/404.html"
	page_404, _ := GetRespBody(uri_404)

	return page_404
}

func Is404(body string)bool{
	if simHtml.GetSimFromStr(vars.Webinfo.Page_404, body)>0.8{
		return true
	}else{
		return false
	}
}