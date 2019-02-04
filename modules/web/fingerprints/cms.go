package fingerprints

import (
	"Shyvana/logger"
	"Shyvana/utils"
	"Shyvana/vars"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

func DetectCms() string{
	var item_s, web_file, match_str, uri, body string
	var idx int
	j_dat, err := ioutil.ReadFile("database/dat_cms.txt")
	if err != nil{
		logger.Log.Println("[ Error ][ IOErr ] Load dat_cms.txt Error")
		return ""
	}
	cms_m := make(map[string][]interface{})
	err = json.Unmarshal(j_dat, &cms_m)
	if err != nil{
		logger.Log.Println("[ Error ][ IOErr ] %v", err)
		return ""
	}
	for name, items := range(cms_m){
		for _, item := range(items){
			item_s = fmt.Sprintf("%v", item)
			idx = strings.Index(item_s, " ")
			web_file = item_s[1:idx]
			match_str = item_s[idx+1:len(item_s)-1]
			uri = vars.Webinfo.Web_url+web_file
			body = utils.GetRespBody(uri)
			if strings.Contains(body, "404"){
				continue
			}
			if len(match_str) == 32{
				if utils.MD5(body) == match_str{
					fmt.Println(body)
					fmt.Println(web_file)
					fmt.Println(match_str)
					return name
				}
			}else {
				if strings.Contains(body, match_str){
					fmt.Println(body)
					fmt.Println(web_file)
					fmt.Println(match_str)
					return name
				}
			}
		}
	}
	return ""
}
