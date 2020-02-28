package fingerprints

import (
	"Shyvana/logger"
	"Shyvana/utils"
	"encoding/json"
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

func GetIPInfo(ip string) *IPINFO{
	var ipinfo IPINFO
	uri := "http://ip.taobao.com/service/getIpInfo.php?ip="+ip
	body, status_code := utils.GetRespBody(uri)
	if status_code != 200{
		logger.Log.Println("[ Error ][ GetIPInfoError ] Failed To access taobao api")
		return nil
	}
	if err := json.Unmarshal([]byte(body), &ipinfo); err != nil {
		logger.Log.Println("[ Error ][ GetIPInfoError ] Failed To Parse the IP result")
		return nil
	}
	return &ipinfo
}
