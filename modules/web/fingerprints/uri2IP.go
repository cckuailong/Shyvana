package fingerprints

import (
	"Shyvana/logger"
	"Shyvana/vars"
	"fmt"
	"net"
	"regexp"
)

func GetIPs()[]string{
	domain := GetMainDomain()
	if domain == ""{
		return nil
	}
	ip_l, err := net.LookupHost(domain)
	if err != nil {
		logger.Log.Println(fmt.Sprintf("[ Warning ][ GetIPWarn ] Cannnot Get the IP of %v", domain))
		return nil
	}
	return ip_l
}

func GetMainDomain()string{
	re, _ := regexp.Compile(`(?i)https?://(.*?)/`)
	res := re.FindAllStringSubmatch(vars.Webinfo.Web_url, -1)
	if len(res) == 0{
		logger.Log.Println("[ Warning ][ GetDomainWarn ] Please Check Your Url")
		return ""
	}
	return res[0][1]
}
