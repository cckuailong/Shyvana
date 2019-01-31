package main

import (
	"fmt"
	"regexp"
)

func main(){
	//web.LaunchWebScan()
	key := "PHP-PID"
	match, _ := regexp.MatchString(`PHP\S*`, key)
	fmt.Println(match)
}