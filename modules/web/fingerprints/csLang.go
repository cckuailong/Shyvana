package fingerprints

import (
	"Shyvana/logger"
	"Shyvana/utils"
	"net/http"
	"regexp"
	"strings"
)

func GetCsLang(headers http.Header, body string)(string, error){
	var match bool
	// via Set-Cookie head
	if _, ok:=headers["Set-Cookie"];ok{
		set_cookie := headers["Set-Cookie"][0]
		eq_ind := strings.Index(set_cookie, "=")
		if eq_ind >= 0{
			lang_str := strings.ToLower(set_cookie[:eq_ind])
			switch lang_str {
			case "phpsessid-php":
				return "Php", nil
			case "jsessionid-jsp":
				return "Java", nil
			case "aspsessionid-asp":
				return "Asp", nil
			case "asp.net_sessionid-aspx":
				return "Asp", nil
			default:
				logger.Log.Println("[ Info ][ CsLang ] No CsLang info in Set-Cookie")
			}
		}else{
			logger.Log.Println("[ Info ][ CsLang ] No CsLang info in Set-Cookie")
		}
	}

	for key, _ := range(headers){
		match, _ = regexp.MatchString(`(?i)PHP\S*`, key)
		if match{
			return "Php", nil
		}
		match, _ = regexp.MatchString(`(?i)Java|Servlet|JSP|JBoss|Glassfish|Oracle|JRE|JDK|JSESSIONID`, key)
		if match{
			return "Java", nil
		}
		match, _ = regexp.MatchString(`(?i)ASP.NET|X-AspNet-Version|x-aspnetmvc-version`, key)
		if match{
			return "Asp", nil
		}
		match, _ = regexp.MatchString(`(?i)python|zope|zserver|wsgi|plone|_ZopeId`, key)
		if match{
			return "Python", nil
		}
		match, _ = regexp.MatchString(`(?i)mod_rack|phusion|passenger`, key)
		if match{
			return "Ruby", nil
		}
	}

	return deepMatch(body), nil
}

// Find CS lang in HTML
func deepMatch(body string)string{
	var re *regexp.Regexp
	var res_l []string
	map_l := make(map[string]int)
	// Php
	re, _ = regexp.Compile(`(?i)\.php|\.phtml`)
	res_l = re.FindAllString(body, -1)
	map_l["Php"] = len(res_l)
	// Java
	re, _ = regexp.Compile(`(?i)\.jsp|\.jspx|.do|\.wss|\.action`)
	res_l = re.FindAllString(body, -1)
	map_l["Java"] = len(res_l)
	// Asp
	re, _ = regexp.Compile(`(?i)\.asp|\.aspx|(__VIEWSTATE\W*)`)
	res_l = re.FindAllString(body, -1)
	map_l["Asp"] = len(res_l)
	// Python
	re, _ = regexp.Compile(`(?i)\.py`)
	res_l = re.FindAllString(body, -1)
	map_l["Python"] = len(res_l)
	// Ruby
	re, _ = regexp.Compile(`(?i)\.rb|\.rhtml`)
	res_l = re.FindAllString(body, -1)
	map_l["Ruby"] = len(res_l)
	// Perl
	re, _ = regexp.Compile(`(?i)\.pl|\.cgi`)
	res_l = re.FindAllString(body, -1)
	map_l["Perl"] = len(res_l)

	// Find the Most Counts as the standard of Most possible
	lang, max_cnt := utils.MaxMap(map_l)
	if max_cnt > 2{
		return lang
	}else{
		return "Unknown"
	}
}


