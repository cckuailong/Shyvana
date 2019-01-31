package fingerprints

import (
	"net/http"
)

func GetWebServ(resp_header http.Header) string{
	if _, ok := resp_header["Server"];ok{
		return resp_header["Server"][0]
	}else{
		return ""
	}
}
