package fingerprints

import (
	"Shyvana/utils"
	"net/http"
	"regexp"
)

func GetCsLang(headers http.Header, body string)string{
	var match bool
	for key, _ := range(headers){
		match, _ = regexp.MatchString(`PHP\S*`, key)
		if match{
			return "Php"
		}
		match, _ = regexp.MatchString(`Java|Servlet|JSP|JBoss|Glassfish|Oracle|JRE|JDK|JSESSIONID`, key)
		if match{
			return "Java"
		}
		match, _ = regexp.MatchString(`ASP.NET|X-AspNet-Version|x-aspnetmvc-version`, key)
		if match{
			return "Asp"
		}
		match, _ = regexp.MatchString(`python|zope|zserver|wsgi|plone|_ZopeId`, key)
		if match{
			return "Python"
		}
		match, _ = regexp.MatchString(`mod_rack|phusion|passenger`, key)
		if match{
			return "Ruby"
		}
	}

	return deepMatch(body)
}

// Find CS lang in HTML
func deepMatch(body string)string{
	var re *regexp.Regexp
	var res_l []string
	map_l := make(map[string]int)
	// Php
	re, _ = regexp.Compile(`\.php|\.phtml`)
	res_l = re.FindAllString(body, -1)
	map_l["Php"] = len(res_l)
	// Java
	re, _ = regexp.Compile(`\.jsp|\.jspx|.do|\.wss|\.action`)
	res_l = re.FindAllString(body, -1)
	map_l["Java"] = len(res_l)
	// Asp
	re, _ = regexp.Compile(`\.asp|\.aspx|(__VIEWSTATE\W*)`)
	res_l = re.FindAllString(body, -1)
	map_l["Asp"] = len(res_l)
	// Python
	re, _ = regexp.Compile(`\.py`)
	res_l = re.FindAllString(body, -1)
	map_l["Python"] = len(res_l)
	// Ruby
	re, _ = regexp.Compile(`\.rb|\.rhtml`)
	res_l = re.FindAllString(body, -1)
	map_l["Ruby"] = len(res_l)
	// Perl
	re, _ = regexp.Compile(`\.pl|\.cgi`)
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


