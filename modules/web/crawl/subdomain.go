package crawl

import (
	"Shyvana/logger"
	"Shyvana/modules/web/fingerprints"
	"Shyvana/utils"
)

func GetSubdomains()[]string{
	var subdm string
	var exist bool
	exist_subdm_l := []string{}
	maindm := "." + fingerprints.GetMainDomain()
	// Verify Universal DNS parse
	arbitrary := "UniversalDNSDetect"
	_ , exist = fingerprints.GetOneIP(arbitrary+maindm)
	if exist{
		logger.Log.Println("[ Info ][ SubdmInfo ] The Domain is Universal Parsing")
		return exist_subdm_l
	}

	subdms := utils.LoadFileToList("database/dat_subdomain.txt")
	for _, c := range(subdms){
		subdm = c+maindm
		_, exist = fingerprints.GetOneIP(subdm)
		if exist{
			exist_subdm_l = append(exist_subdm_l, subdm)
		}
	}
	return exist_subdm_l
}
