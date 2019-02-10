package fingerprints

import (
	"Shyvana/logger"
	"github.com/likexian/whois-go"
	"github.com/likexian/whois-parser-go"
)

func GetWhoisInfo()(*whois_parser.WhoisInfo, error){
	uri := GetMainDomain()
	res, err0 := whois.Whois(uri)
	if err0 != nil{
		logger.Log.Printf("[ Error ][ WhoisErr ] Get Raw Whois Info Error(%v)\n", err0)
		return nil, err0
	}
	ws, err1 := whois_parser.Parse(res)
	if err1 != nil{
		logger.Log.Printf("[ Error ][ WhoisErr ] Parse Raw Whois Info Error(%v)\n", err1)
		return nil, err0
	}
	return &ws, nil
}
