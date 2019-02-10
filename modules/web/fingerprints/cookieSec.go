package fingerprints

import (
	"regexp"
)

type COOKIERES struct {
	secure bool
	httponly bool
	domain []string
	path []string
}

func DetectCookieSec(cookies string) (*COOKIERES, error){
	var match bool
	var err error
	var re *regexp.Regexp
	var tmp_l [][]string
	cookieres := COOKIERES{
		secure:false,
		httponly:false,
		domain:[]string{},
		path:[]string{},
	}
	// secure flag
	match, err = regexp.MatchString(`(?i)secure;`, cookies)
	if err != nil{
		return nil, err
	}
	if match{
		cookieres.secure = true
	}
	// httponly flag
	match, err = regexp.MatchString(`(?i)HttpOnly`, cookies)
	if err != nil{
		return nil, err
	}
	if match{
		cookieres.httponly = true
	}
	// domain
	re, _ = regexp.Compile(`(?i)domain\=(.+?);`)
	tmp_l = re.FindAllStringSubmatch(cookies, -1)
	domain_l := []string{}
	for _, c := range(tmp_l){
		domain_l = append(domain_l, c[1])
	}
	if len(domain_l) != 0{
		cookieres.domain = domain_l
	}
	// path
	re, _ = regexp.Compile(`(?i)path\=(.*?);`)
	tmp_l = re.FindAllStringSubmatch(cookies, -1)
	path_l := []string{}
	for _, c := range(tmp_l){
		path_l = append(path_l, c[1])
	}
	if len(path_l) != 0{
		cookieres.path = path_l
	}
	return &cookieres, nil
}
