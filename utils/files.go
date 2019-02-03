package utils

import (
	"Shyvana/logger"
	"io/ioutil"
	"strings"
)

func LoadFileToList(filename string)[]string{
	content, err := ioutil.ReadFile(filename)
	if err != nil{
		logger.Log.Println("[ Error ][ IOErr ] %v", err)
		return nil
	}
	body := string(content)
	body = strings.TrimSpace(body)
	res_l := strings.Split(body, "\n")
	return res_l
}
