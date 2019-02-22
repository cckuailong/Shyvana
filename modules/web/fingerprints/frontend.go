package fingerprints

import (
	"Shyvana/logger"
	"github.com/buger/jsonparser"
	"io/ioutil"
	"regexp"
)

func DetectFrontEnd(body string)(string, error){
	detected := "Unknown"
	j_dat, err := ioutil.ReadFile("database/dat_frontend.txt")
	if err != nil{
		logger.Log.Println("[ Error ][ IOErr ] Load dat_frontend.txt Error")
		return "", err
	}
	err = jsonparser.ObjectEach(j_dat, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		_,err = jsonparser.ArrayEach(value, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			//fmt.Println(string(value))
			match, _ := regexp.MatchString(`(?i)`+string(value), body)
			if match{
				detected = string(key)
				return
			}
		}, "body")
		return nil
	})
	if err != nil{
		logger.Log.Printf("[ Error ][ JsonErr ] Parse JsonFile dat_frontend.txt Error(%v)\n", err)
		return "", err
	}
	return detected, nil
}
