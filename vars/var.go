package vars

import (
	"Shyvana/settings"
)

type (
		WEBINFO struct {
			Web_url  string
		}
	)

var(
	Webinfo WEBINFO
	Headers map[string]string
)

func init(){
	web_info := settings.Cfg.Section("WEB")
	Webinfo.Web_url = web_info.Key("WEB_URL").MustString("http://l0vebear.top/wordpress/")

	Headers = make(map[string]string)
	Headers["Accept"] = "application/ "
	Headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36"

}