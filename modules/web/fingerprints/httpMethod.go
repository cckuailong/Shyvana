package fingerprints

import (
	"Shyvana/logger"
	"Shyvana/utils"
	"net/http"
)

func GetHttpMethod(resp_opt_header http.Header)[]string{

	if _, ok := resp_opt_header["Allow"];ok{
		allowed := resp_opt_header["Allow"]
		if utils.StrIsIn("PUT", allowed) || utils.StrIsIn("DELETE", allowed){
			logger.Log.Println("[Warning][ OptWarn ] Dangerous Http Method Found")
		}
		return allowed
	}else{
		return []string{}
	}
}
