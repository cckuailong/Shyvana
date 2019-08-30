package vul

import (
	"fmt"
)


func RunSqlmap(crawled_l []string){
	for _, uri := range(crawled_l){
		fmt.Println(uri)
	}

}