package vul

import (
	"Shyvana/logger"
	"Shyvana/utils"
	"Shyvana/vars"
	"bytes"
	"errors"
	"fmt"
	"github.com/buger/jsonparser"
	"io/ioutil"
	"net/http"
	"os/exec"
	"regexp"
	"time"
)

type SQLTASK struct {
	Taskid     string `json:"taskid" db:",omitempty,json"`
	Url        string `json:"url" db:",json"`
	Engineid   string `json:"engineid",db:",json"`
	Status     string `json:"status" db:",json"`
	Vul        bool   `json:"vul" db:",json"`
	VulInfo    map[string][]string
}

func GetVulInfo(data []byte)(map[string][]string, error){
	vulinfo := map[string][]string{}
	_,err := jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		type_n,_,_,_ := jsonparser.Get(value, "type")
		if string(type_n) == "0"{
			scan_val,_,_,_ := jsonparser.Get(value, "value")
			scan_uri,_,_,_ := jsonparser.Get(scan_val, "url")
			scan_uri_l := []string{string(scan_uri)}
			vulinfo["url"] = scan_uri_l
		}else{
			_,err = jsonparser.ArrayEach(value, func(value []byte, dataType jsonparser.ValueType, offset int, err error){
				// db
				dbms,_,_,_ := jsonparser.Get(value, "dbms")
				db_v,_,_,_ := jsonparser.Get(value, "dbms_version", "[0]")
				vulinfo["dbms"] = []string{string(dbms) + string(db_v)}
				// parameter
				param,_,_,_ := jsonparser.Get(value, "parameter")
				params, ok := vulinfo["params"]
				if ok{
					params = append(params, string(param))
					vulinfo["params"] = params
				}else{
					vulinfo["params"] = []string{string(param)}
				}
				vuls,_,_,_ := jsonparser.Get(value, "data")
				err = jsonparser.ObjectEach(vuls, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
					//vul title
					title,_,_,_ := jsonparser.Get(value, "title")
					titles, ok := vulinfo["titles"]
					if ok{
						titles = append(titles, string(title))
						vulinfo["titles"] = titles
					}else{
						vulinfo["titles"] = []string{string(title)}
					}
					// payloads
					payload,_,_,_ := jsonparser.Get(value, "payload")
					payloads, ok := vulinfo["payloads"]
					if ok{
						payloads = append(payloads, string(payload))
						vulinfo["payloads"] = payloads
					}else{
						vulinfo["payloads"] = []string{string(payload)}
					}
					return nil
				})
			}, "value")
		}
	}, "data")
	if err != nil{
		logger.Log.Printf("[ Error ][ JsonErr ] Parse Json Error(%v)\n", err)
		return nil, err
	}
	return vulinfo, nil
}

func (sqltask *SQLTASK)NewTask()error{
	body := utils.GetRespBody(vars.Webinfo.Local+"task/new")
	flag,datatype,_,err := jsonparser.Get([]byte(body), "success")
	if datatype == jsonparser.NotExist{
		logger.Log.Println("[ Error ][ SQlErr ] SQLmap Start Error")
		return err
	}
	if string(flag) == "true"{
		taskid,datatype,_,err := jsonparser.Get([]byte(body), "taskid")
		if datatype == jsonparser.NotExist{
			logger.Log.Println("[ Error ][ SQlErr ] SQLmap Start Error")
			return err
		}else{
			sqltask.Taskid = string(taskid)
		}
	}else{
		logger.Log.Println("[ Error ][ SQlErr ] SQLmap Start Error")
		return errors.New("New Task Create Failed")
	}
	return nil
}

func (sqltask *SQLTASK)DelTask()error{
	body := utils.GetRespBody(vars.Webinfo.Local + "task/" + sqltask.Taskid + "/delete")
	flag,datatype,_,err := jsonparser.Get([]byte(body), "success")
	if datatype == jsonparser.NotExist{
		logger.Log.Println("[ Error ][ SQlErr ] SQLmap Start Error")
		return err
	}else{
		if string(flag) == "true"{
			return nil
		}
	}
	return errors.New("Task Delete Failed")
}

func (sqltask *SQLTASK)StartScan()error{
	uri := vars.Webinfo.Local + "scan/" + sqltask.Taskid + "/start"
	vals := `{"url": "`+sqltask.Url+`"}`
	pv := bytes.NewBuffer([]byte(vals))
	req, err := http.NewRequest("POST", uri, pv)
	if err != nil{
		logger.Log.Println("%v", err)
		return err
	}
	req.Header.Set("Content-type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		return err
	}
	body,  _ := ioutil.ReadAll(resp.Body)
	flag,datatype,_,err := jsonparser.Get(body, "success")
	if datatype == jsonparser.NotExist{
		logger.Log.Println("[ Error ][ SQlErr ] SQLmap Start Scan Error")
		return err
	}
	if string(flag) == "true"{
		engineid,datatype,_,err := jsonparser.Get(body, "engineid")
		if datatype == jsonparser.NotExist{
			logger.Log.Println("[ Error ][ SQlErr ] SQLmap Start Scan Error")
			return err
		}else{
			sqltask.Engineid = string(engineid)
		}
	}else{
		logger.Log.Println("[ Error ][ SQlErr ] SQLmap Start Error")
		return errors.New("Start Scan Failed")
	}

	return nil
}

func (sqltask *SQLTASK)ScanStatus()error{
	body := utils.GetRespBody(vars.Webinfo.Local+ "scan/" + sqltask.Taskid + "/status")
	flag,datatype,_,err := jsonparser.Get([]byte(body), "success")
	if datatype == jsonparser.NotExist{
		logger.Log.Println("[ Error ][ SQlErr ] SQLmap Start Scan Error")
		return err
	}
	if string(flag) == "true"{
		status,datatype,_,err := jsonparser.Get([]byte(body), "status")
		if datatype == jsonparser.NotExist{
			logger.Log.Println("[ Error ][ SQlErr ] SQLmap Start Scan Error")
			return err
		}else{
			sqltask.Status = string(status)
		}
	}else{
		logger.Log.Println("[ Error ][ SQlErr ] SQLmap Start Error")
		return errors.New("Get Scan Status Failed")
	}
	return nil
}

func (sqltask *SQLTASK)ScanData()error{
	body := utils.GetRespBody(vars.Webinfo.Local+"scan/" + sqltask.Taskid + "/data")
	data,datatype,_,err := jsonparser.Get([]byte(body), "data")
	if datatype == jsonparser.NotExist{
		logger.Log.Println("[ Error ][ SQlErr ] SQLmap Scan Data Error")
		return err
	}else{
		if len(data) > 0{
			sqltask.Vul = true
			sqltask.VulInfo, err = GetVulInfo([]byte(body))
			if err != nil{
				return err
			}
		}else{
			sqltask.Vul = false
			sqltask.VulInfo = nil
		}
	}
	return nil
}

func (sqltask *SQLTASK)SetOptions()error{
	uri := vars.Webinfo.Local + "option/" + sqltask.Taskid + "/set"
	vals := `{"options": {"randomAgent": "True", "tech":"BT"}}`
	pv := bytes.NewBuffer([]byte(vals))
	req, err := http.NewRequest("POST", uri, pv)
	if err != nil{
		logger.Log.Println("%v", err)
		return err
	}
	req.Header.Set("Content-type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		return err
	}
	body,  _ := ioutil.ReadAll(resp.Body)
	flag,datatype,_,err := jsonparser.Get([]byte(body), "success")
	if datatype == jsonparser.NotExist{
		logger.Log.Println("[ Error ][ SQlErr ] SQLmap Start Scan Error")
		return err
	}
	if string(flag) != "true"{
		logger.Log.Println("[ Error ][ SQlErr ] SQLmap Start Error")
		return errors.New("Set Task Options Failed")
	}
	return nil
}

func (sqltask *SQLTASK)StopScan()error{
	body := utils.GetRespBody(vars.Webinfo.Local + "scan/" + sqltask.Taskid + "/stop")
	flag,datatype,_,err := jsonparser.Get([]byte(body), "success")
	if datatype == jsonparser.NotExist{
		logger.Log.Println("[ Error ][ SQlErr ] SQLmap Start Scan Error")
		return err
	}
	if string(flag) != "true"{
		logger.Log.Println("[ Error ][ SQlErr ] SQLmap Start Error")
		return errors.New("Stop Task Failed")
	}
	return nil
}

func (sqltask *SQLTASK)KillScan()error{
	body := utils.GetRespBody(vars.Webinfo.Local + "scan/" + sqltask.Taskid + "/kill")
	flag,datatype,_,err := jsonparser.Get([]byte(body), "success")
	if datatype == jsonparser.NotExist{
		logger.Log.Println("[ Error ][ SQlErr ] SQLmap Start Scan Error")
		return err
	}
	if string(flag) != "true"{
		logger.Log.Println("[ Error ][ SQlErr ] SQLmap Start Error")
		return errors.New("Kill Task Failed")
	}
	return nil
}

func (sqltask *SQLTASK)RunSingle(){
	err := sqltask.NewTask()
	if err != nil{
		logger.Log.Printf("[ Error ][ SQLErr ] %v\n", err)
		return
	}
	err = sqltask.SetOptions()
	if err != nil{
		logger.Log.Printf("[ Error ][ SQLErr ] %v\n", err)
		return
	}
	err = sqltask.StartScan()
	if err != nil{
		logger.Log.Printf("[ Error ][ SQLErr ] %v\n", err)
		return
	}
	start_time := time.Now()
	flag := true
	for flag{
		err = sqltask.ScanStatus()
		if err != nil{
			logger.Log.Printf("[ Error ][ SQLErr ] %v\n", err)
			return
		}
		switch sqltask.Status {
		case "terminated":
			flag = false
			break
		case "running":
			time.Sleep(10 * time.Second)
			break
		default:
			flag = false
			break
		}
		if time.Now().Sub(start_time) > 500*time.Second {
			err = sqltask.StopScan()
			if err != nil{
				err = sqltask.KillScan()
				if err != nil{
					logger.Log.Printf("[ Error ][ SQLErr ] %v\n", err)
					return
				}
			}
			time.Sleep(time.Second)
			flag = false
			break
		}
	}
	err = sqltask.ScanData()
	if err != nil{
		logger.Log.Printf("[ Error ][ SQLErr ] %v\n", err)
		return
	}
}

func RunSqlmap(crawled_l []string){
	cmd := exec.Command("python", "libs/sqlmap/sqlmapapi.py", "-s")
	cmd.Start()
	time.Sleep(2*time.Second)

	for _, uri := range(crawled_l){
		match, _ := regexp.MatchString(`(.*?)\?(.*?)=(.*?)`, uri)
		if match{
			sqltask := new(SQLTASK)
			sqltask.RunSingle()
			fmt.Println(sqltask.VulInfo)
		}
	}

	fmt.Println(cmd.Process.Pid)
	cmd.Process.Kill()
}