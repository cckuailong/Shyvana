package crawl

import (
	"Shyvana/logger"
	"Shyvana/modules/web/net"
	"Shyvana/utils"
	"Shyvana/vars"
)

func GetSubdomains()[]string{
	var subdm string
	var exist bool
	exist_subdm_l := []string{}
	maindm := "." + net.GetMainDomain(vars.Webinfo.Web_url)
	// Verify Universal DNS parse
	arbitrary := "UniversalDNSDetect"
	_ , exist = net.GetOneIP(arbitrary+maindm)
	if exist{
		logger.Log.Println("[ Info ][ SubdmInfo ] The Domain is Universal Parsing")
		return exist_subdm_l
	}

	subdms := utils.LoadFileToList("database/dat_subdomain.txt")
	for _, c := range(subdms){
		subdm = c+maindm
		_, exist = net.GetOneIP(subdm)
		if exist{
			exist_subdm_l = append(exist_subdm_l, subdm)
		}
	}
	return exist_subdm_l
}
