package fingerprints

import (
	"Shyvana/logger"
	"Shyvana/utils"
	"github.com/buger/jsonparser"
	"io/ioutil"
	"net/http"
)

func DetectCdn(header http.Header)(string, error){
	detected := "Unknown"
	j_dat, err := ioutil.ReadFile("database/dat_cdn.txt")
	if err != nil{
		logger.Log.Println("[ Error ][ IOErr ] Load dat_cdn.txt Error")
		return "", err
	}
	err = jsonparser.ObjectEach(j_dat, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		categ,datatype,_,_ := jsonparser.Get(value, "headers")
		if datatype != jsonparser.NotExist{
			err = jsonparser.ObjectEach(categ, func(key1 []byte, value1 []byte, dataType jsonparser.ValueType, offset int) error {
				if _,ok := header[string(key1)];ok{
					_,err = jsonparser.ArrayEach(categ, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
						if utils.StrLikelyIn(string(value), header[string(key1)]){
							detected = string(key)
							return
						}
					}, string(key1))
				}
				return nil
			})
		}
		return nil
	})
	if err != nil{
		logger.Log.Printf("[ Error ][ JsonErr ] Parse JsonFile dat_cdn.txt Error(%v)\n", err)
		return "", err
	}
	return detected, nil
}
