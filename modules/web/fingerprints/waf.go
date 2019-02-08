package fingerprints

import (
	"Shyvana/logger"
	"Shyvana/utils"
	"github.com/buger/jsonparser"
	"io/ioutil"
	"net/http"
	"strings"
)

func DetectWaf(header http.Header, body string)string{
	detected := "Unknown"
	j_dat, err := ioutil.ReadFile("database/dat_waf.txt")
	if err != nil{
		logger.Log.Println("[ Error ][ IOErr ] Load dat_cms.txt Error")
		return ""
	}
	err = jsonparser.ObjectEach(j_dat, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		categ,datatype,_,_ := jsonparser.Get(value, "index")
		if datatype != jsonparser.NotExist{
			_,err := jsonparser.ArrayEach(categ, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				if strings.Contains(body, string(value)){
					detected = string(key)
					return
				}
			}, "index")
			if err != nil{
				return err
			}
		}

		categ,datatype,_,_ = jsonparser.Get(value, "headers")
		if dataType != jsonparser.NotExist{
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
		logger.Log.Printf("[ Error ][ JsonErr ] Parse JsonFile dat_waf.txt Error(%v)\n", err)
		return ""
	}
	return ""
}
