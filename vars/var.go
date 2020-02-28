package vars

import (
	"Shyvana/modules/web/net"
	"Shyvana/settings"
	"strings"
)

type (
		WEBINFO struct {
			Input_url string
			Web_url string
			Web_domain string
			Page_404 string
			Local string
		}
	)

var(
	Webinfo WEBINFO
	Headers map[string]string
)

func init(){
	var proto string
	web_info := settings.Cfg.Section("WEB")
	Webinfo.Input_url = web_info.Key("WEB_URL").MustString("http://lovebear.top/wordpress/")
	if strings.HasPrefix(Webinfo.Input_url, "https"){
		proto = "https://"
	}else{
		proto = "http://"
	}
	Webinfo.Web_domain = net.GetMainDomain(Webinfo.Input_url)
	Webinfo.Web_url = proto+Webinfo.Web_domain

	Webinfo.Local = web_info.Key("LOCAL").MustString("http://127.0.0.1:8775/")

	Headers = make(map[string]string)
	Headers["Accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3"
	Headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36"

}