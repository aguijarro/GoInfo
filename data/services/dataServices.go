package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Data struct {
	Title string `json:"title"`
	Description string `json:"description"`
	User string `json:"user"`
	Result [][]string `json:"result"`
}

type Result struct{
	Description string
	ValYear1 string
	ValYear2 string
	ValYear3 string
	ValYear4 string
	ValYear5 string		
	ValYear6 string			
}

func ShowContent(url string, initSet int, endSet int)(results []Result, err error){

		
		data, _ := makeContent(url)
		dataSet := data.Result[initSet:endSet]
	
		
	for _,data1 := range dataSet{

		result := Result{
			Description : data1[0],
			ValYear1 : data1[1],
			ValYear2 : data1[2],
			ValYear3 : data1[3],
			ValYear4 : data1[4],
			ValYear5 : data1[5],	 
			ValYear6 : data1[6],	
		}	
			
		results = append(results, result)
	}		
	
	return
}

func readServices(url string)([]byte, error){

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	
	return body, nil
	
}

func makeContent(url string)(*Data, error){
	content, err := readServices(url)
	if err != nil {
		return nil, err
	}
	
	var data Data
	
	err = json.Unmarshal(content, &data)
	
	if err != nil{
		return nil, err
	}
	
	return &data, err
}

