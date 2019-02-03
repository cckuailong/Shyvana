package fingerprints

import (
	"Shyvana/logger"
	"Shyvana/utils"
	"Shyvana/vars"
	"encoding/json"
	"fmt"
	"net"
	"regexp"
	"strings"
)

type IPINFO struct {
	Code int
	Data IP
}

type IP struct {
	Country   string `json:"country"`
	CountryId string `json:"country_id"`
	Area      string `json:"area"`
	AreaId    string `json:"area_id"`
	Region    string `json:"region"`
	RegionId  string `json:"region_id"`
	City      string `json:"city"`
	CityId    string `json:"city_id"`
	Isp       string `json:"isp"`
}

func GetOneIP(domain string)(string, bool){
	ipAddr, err := net.ResolveIPAddr("ip", domain)
	if err != nil{
		return "", false
	}
	res := ipAddr.String()
	if strings.Contains(res, "no such host"){
		return "", false
	}
	return res, true
}

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
	re, _ := regexp.Compile(`(?i)http[s]?://[www]?(.*?)/`)
	res := re.FindAllStringSubmatch(vars.Webinfo.Web_url, -1)
	if len(res) == 0{
		logger.Log.Println("[ Warning ][ GetDomainWarn ] Please Check Your Url")
		return ""
	}
	return res[0][1]
}

func GetIPInfo(ip string) *IPINFO{
	var ipinfo IPINFO
	uri := "http://ip.taobao.com/service/getIpInfo.php?ip="+ip
	body := utils.GetRespBody(uri)
	if err := json.Unmarshal([]byte(body), &ipinfo); err != nil {
		return nil
	}
	return &ipinfo
}