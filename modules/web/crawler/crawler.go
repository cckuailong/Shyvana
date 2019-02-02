package crawler

import (
	"Shyvana/modules/web/fingerprints"
	"Shyvana/utils"
	"Shyvana/vars"
	"container/list"
	"fmt"
	"regexp"
	"strings"
)


func Crawl(ip_l []string)[]string{
	var body string
	crawling_l := list.New()
	crawled_l := []string{}
	maindom := fingerprints.GetMainDomain()

	crawling_l.PushBack(vars.Webinfo.Web_url)
	for crawling_l.Len() != 0{
		front := crawling_l.Front()
		uri := fmt.Sprint(front.Value)
		if utils.StrIsIn(uri, crawled_l){
			crawling_l.Remove(front)
			continue
		}
		body = utils.GetRespBody(uri)
		re, _ := regexp.Compile(`(?i)href\s?=\s?["\|'](.*?)["\|']`)
		uris := re.FindAllStringSubmatch(body, -1)
		for _, uri := range(uris){
			if filterUri(uri[1], maindom, ip_l[0]){
				crawling_l.PushBack(uri[1])
				//fmt.Println(uri[1])
			}
		}
		crawled_l = append(crawled_l, uri)
		crawling_l.Remove(front)
	}
	return crawled_l
}

func filterUri(uri, maindom, ip string)bool{
	if (strings.Contains(uri, maindom) || strings.Contains(uri, ip)) && strings.Contains(uri, "http"){
		return true
	}else{
		return false
	}
}
