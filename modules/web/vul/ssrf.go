package vul

import (
	"Shyvana/logger"
	"Shyvana/modules/web/fingerprints"
	"Shyvana/utils"
	"strconv"
	"strings"
)

func ip2long(ip string)uint32{
	var res uint32 = 0
	var bit_len uint32 = 24
	tmp_l := strings.Split(ip, ".")
	if len(tmp_l) > 4{
		logger.Log.Println("[ Error ][ SSRFError ] Illegal IP Address String")
	}
	for _, tmp :=range(tmp_l[:len(tmp_l)-1]){
		num,err := strconv.ParseUint(tmp,0, 32)
		if err != nil{
			panic(err)
		}
		res += uint32(num<<bit_len)
		bit_len -= 8
	}
	last_num,err := strconv.ParseUint(tmp_l[len(tmp_l)-1], 0, 32)
	if err != nil{
		panic(err)
	}
	return res+uint32(last_num)
}

func is_inner_addr(ip string) bool{
	ip_long := ip2long(ip)
	return ip2long("127.0.0.0")>>24 == ip_long>>24 ||
		ip2long("10.0.0.0")>>24 == ip_long>>24 ||
		ip2long("172.16.0.0")>>20 == ip_long>>20 ||
		ip2long("192.168.0.0")>>16 == ip_long>>16
}

func check_inner(uri string) bool{
	domain := fingerprints.GetPureUri(uri)
	ip,ok := fingerprints.GetOneIP(domain)
	if !ok{
		logger.Log.Println("[ Error ][ SSRFError ] Failed To Translate Domain To IP")
		return false
	}
	return is_inner_addr(ip)
}

func Check_ssrf(uri string) bool{
	if check_inner(uri){
		return true
	}
	headers,status_code := utils.GetRespHeaderNoRedirect(uri)
	for status_code-300>=0 && status_code-300<10{
		redirect_uri := headers["Location"][0]
		if strings.HasPrefix(redirect_uri, "//"){  // //ip:port/mid/test.html
			redirect_uri = "http"+redirect_uri
		}else if strings.HasPrefix(redirect_uri, "http"){ // http(s)://ip:port/mid/test.html
			redirect_uri = redirect_uri
		} else if strings.HasPrefix(redirect_uri, "/"){ // /test.html
			redirect_uri = fingerprints.GetPureUri(uri)+redirect_uri
		}else{ // test.html
			redirect_uri = uri+redirect_uri
		}
		if check_inner(redirect_uri){
			return true
		}
		headers,status_code = utils.GetRespHeaderNoRedirect(redirect_uri)
	}
	return false
}