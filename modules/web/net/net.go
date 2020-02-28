package net

import (
	"Shyvana/logger"
	"fmt"
	"net"
	"regexp"
	"strings"
)


func GetOneIP(domain string)(string, bool){
	ipAddr, err := net.ResolveIPAddr("ip", domain)
	if err != nil{
		return "", false
	}
	res := ipAddr.String()
	if strings.Contains(res, "no such host"){
		return "", false
	}
	if res == ""{
		return "",false
	}
	return res, true
}

func GetIPs(uri string)[]string{
	domain := GetMainDomain(uri)
	fmt.Println(domain)
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

func GetMainDomain(uri string)string{
	re, _ := regexp.Compile(`(?i)^http[s]?://[www\.]?(.*?)[/]*$`)
	res := re.FindAllStringSubmatch(uri, -1)
	if len(res) == 0{
		logger.Log.Println("[ Warning ][ GetDomainWarn ] Please Check Your Url")
		return ""
	}

	return res[0][len(res[0])-1]
}
