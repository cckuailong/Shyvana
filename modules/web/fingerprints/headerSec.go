package fingerprints

import (
	"Shyvana/utils"
	"net/http"
)

type HEADERRES struct {
	x_xss bool
	x_frame bool
	content_type bool
	strict_trans bool
	x_content_type bool
	strange_header []string
}

var common_header []string

func DetectHeaderSec(headers http.Header)*HEADERRES{
	headerres := HEADERRES{
		x_xss:false,
		x_frame:false,
		content_type:false,
		strict_trans:false,
		x_content_type:false,
		strange_header:[]string{},
	}
	if _,ok:=headers["X-XSS-Protection"];ok{
		headerres.x_xss = true
	}
	if _,ok:=headers["X-Frame-Options"];ok{
		headerres.x_frame = true
	}
	if _,ok:=headers["Content-Type"];ok{
		headerres.content_type = true
	}
	if _,ok:=headers["Strict-Transport-Security"];ok{
		headerres.strict_trans = true
	}
	if _,ok:=headers["X-Content-Type-Options"];ok{
		headerres.x_content_type = true
	}
	strange_l := []string{}
	for key, _ := range(headers){
		if !utils.StrIsIn(key, common_header){
			strange_l = append(strange_l, key)
		}
	}
	if len(strange_l) != 0{
		headerres.strange_header = strange_l
	}
	return &headerres
}

func init(){
	common_header = []string{"Accept", "Accept-Charset", "Accept-Encoding", "Accept-Language",
		"Accept-Datetime", "Authorization", "Connection", "Cookie", "Content-Length", "Content-MD5",
		"Content-Type", "Expect", "From", "Host", "If-Match", "If-Modified-Since", "If-None-Match",
		"If-Range", "If-Unmodified-Since", "Max-Forwards", "Origin", "Pragma", "Proxy-Authorization",
		"Range", "Referer", "User-Agent", "Upgrade", "Via", "Warning", "X-Requested-With",
		"X-Forwarded-For", "X-Forwarded-Host", "X-Forwarded-Proto", "Front-End-Https",
		"X-Http-Method-Override", "X-ATT-DeviceId", "X-Wap-Profile", "Proxy-Connection",
		"Accept-Ranges", "Age", "Allow", "Cache-Control", "Content-Encoding", "Content-Language",
		"Content-Length", "Content-Location", "Content-MD5", "Content-Disposition", "Content-Range",
		"Content-Type", "Date", "ETag", "Expires", "Last-Modified", "Link", "Location",
		"Proxy-Authenticate", "Refresh", "Retry-After", "Server", "Set-Cookie", "Status",
		"Strict-Transport-Security", "Trailer", "Transfer-Encoding", "Vary", "WWW-Authenticate",
		"X-Frame-Options", "Public-Key-Pins", "X-XSS-Protection", "Content-Security-Policy",
		"X-Content-Security-Policy", "X-WebKit-CSP", "X-Content-Type-Options", "X-Powered-By",
		"Keep-Alive", "Content-language", "X-UA-Compatible"}
}
