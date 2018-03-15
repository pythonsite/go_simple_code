//测试客户端代码
package main

import (
	"bytes"
	"go_dev/11/short_url/model"
	"encoding/json"
	"net/http"
	"fmt"
	"io/ioutil"
)

func getShortUrl()string{
	// 根据端地址获取长地址
	var buffer bytes.Buffer
	var req model.Long2ShortRequest
	req.OriginUrl="https://www.py2222ite.com/p/23211123123123323?utm_campaign=hugo&utm_medium=reader_share&utm_content=note&utm_source=qq"
	data,_ := json.Marshal(req)
	buffer.WriteString(string(data))
	resp,err := http.Post("http://127.0.0.1:18888/trans/long2short","application/json",&buffer)
	if err != nil{
		fmt.Println("post failed,err:",err)
		return ""
	}
	result,_:= ioutil.ReadAll(resp.Body)
	var response model.Long2ShortResponse
	err = json.Unmarshal(result,&response)
	res := "http://127.0.0.1:8888/"+response.ShortUrl
	return res
}

func getLongUrl()string{
	// 根据长地址获取端地址
	var buffer bytes.Buffer
	var req model.Short2LongRequest
	req.ShortUrl="4c95"
	data,_ := json.Marshal(req)
	buffer.WriteString(string(data))
	resp,err := http.Post("http://127.0.0.1:18888/trans/short2long","application/json",&buffer)
	if err != nil{
		fmt.Println("post failed,err:",err)
		return ""
	}
	result,_:= ioutil.ReadAll(resp.Body)
	var response model.Short2LongResponse
	err = json.Unmarshal(result,&response)

	res := response.OriginUrl
	return res
}

func main() {
	short_url := getShortUrl()
	fmt.Printf("short url:%v\n",short_url)
	origin_url := getLongUrl()
	fmt.Printf("origin url:%v\n",origin_url)


}
