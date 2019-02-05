package fingerprints

import (
	"Shyvana/logger"
	"fmt"
	"github.com/buger/jsonparser"
	"io/ioutil"
)

func DetectWaf()string{
	j_dat, err := ioutil.ReadFile("database/dat_waf.txt")
	if err != nil{
		logger.Log.Println("[ Error ][ IOErr ] Load dat_cms.txt Error")
		return ""
	}
	err = jsonparser.ObjectEach(j_dat, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		categ,datatype,_,_ := jsonparser.Get(value, "index")
		if datatype != jsonparser.NotExist{
			_,err := jsonparser.ArrayEach(categ, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				fmt.Println(string(value))
			}, "index")
			if err != nil{
				return err
			}
		}
		categ,datatype,_,_ = jsonparser.Get(value, "headers")
		return nil
	})
	if err != nil{
		logger.Log.Printf("[ Error ][ JsonErr ] Parse JsonFile dat_waf.txt Error(%v)\n", err)
		return ""
	}
	return ""
}
