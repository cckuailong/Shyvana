package fingerprints

import (
	"net/http"
	"regexp"
)

func isPhp(resp http.Response)bool{
	for key, _ := range(resp.Header){
		match, _ := regexp.MatchString(`PHP\S*`, key)
		return true
	}
}
