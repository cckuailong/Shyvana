package fingerprints

import (
	"Shyvana/logger"
	"Shyvana/modules/web/net"
	"Shyvana/vars"
	"github.com/likexian/whois-go"
	"github.com/likexian/whois-parser-go"
)

func GetWhoisInfo()(*whoisparser.WhoisInfo, error){
	uri := net.GetMainDomain(vars.Webinfo.Web_url)
	res, err0 := whois.Whois(uri)
	if err0 != nil{
		logger.Log.Printf("[ Error ][ WhoisErr ] Get Raw Whois Info Error(%v)\n", err0)
		return nil, err0
	}
	ws, err1 := whoisparser.Parse(res)
	if err1 != nil{
		logger.Log.Printf("[ Error ][ WhoisErr ] Parse Raw Whois Info Error(%v)\n", err1)
		return nil, err0
	}
	return &ws, nil
}
