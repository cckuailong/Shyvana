package crawl

import (
	"Shyvana/modules/web/net"
	"Shyvana/utils"
	"Shyvana/vars"
	"container/list"
	"fmt"
	"regexp"
	"strings"
)

// Input (uri's different IP , use to filter)
// return (Website URIs, Found Emails)
func Crawl()([]string, []string){
	fetched_email := []string{}
	crawling_l := list.New()
	crawled_l := []string{}
	maindom := net.GetMainDomain(vars.Webinfo.Web_url)

	crawling_l.PushBack(vars.Webinfo.Web_url)
	for crawling_l.Len() != 0{
		front := crawling_l.Front()
		cur_uri := fmt.Sprint(front.Value)
		if utils.StrIsIn(cur_uri, crawled_l){
			crawling_l.Remove(front)
			continue
		}
		body, status_code := utils.GetRespBody(cur_uri)
		if status_code == 404 || utils.Is404(body){
			crawling_l.Remove(front)
			continue
		}
		fetched_email = FetchEmail(body, fetched_email)
		re, _ := regexp.Compile(`(?i)href\s?=\s?["\|']#?(.*?)["\|']`)
		uris := re.FindAllStringSubmatch(body, -1)
		for _, uri := range(uris){
			uu := uri[1]
			if strings.HasPrefix(uu, "http"){
				if filterUri(uu, maindom){
					crawling_l.PushBack(uu)
				}
			}else{
				uu = cur_uri + uu
				crawling_l.PushBack(uu)
			}
		}

		crawled_l = append(crawled_l, cur_uri)
		crawling_l.Remove(front)
	}
	return crawled_l, fetched_email
}

// Filter the Appropriate Uri
func filterUri(uri, maindom string)bool{
	if strings.Contains(uri, maindom) && strings.Contains(uri, "http"){
		return true
	}else{
		return false
	}
}
