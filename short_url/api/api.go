package main

import (
	"io/ioutil"
	"net/http"
	"fmt"
	"encoding/json"
	"go_dev/11/short_url/logic"
	"go_dev/11/short_url/model"
	_ "github.com/go-sql-driver/mysql"
)

const (
	ErrSuccess = 0
	ErrInvalidParameter = 1001
	ErrServerBusy = 1002
)

func getMessage(code int) (msg string){
	switch code {
	case ErrSuccess:
		msg = "success"
	case ErrInvalidParameter:
		msg = "invalid parameter"
	case ErrServerBusy:
		msg = "server busy"
	default:
		msg = "unknown error"
	}

	return
}

// 用于将返回序列化数据，失败的返回
func responseError(w http.ResponseWriter, code int) {
	var response model.ResponseHeader
	response.Code = code
	response.Message = getMessage(code)

	data, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("{\"code\":500, \"message\": \"server busy\"}"))
		return
	}

	w.Write(data)
}

// 用于将返回序列化数据，成功的返回
func responseSuccess(w http.ResponseWriter, data interface{}) {


	dataByte, err := json.Marshal(data)
	if err != nil {
		w.Write([]byte("{\"code\":500, \"message\": \"server busy\"}"))
		return
	}

	w.Write(dataByte)
}

// 长地址到短地址
func Long2Short(w http.ResponseWriter, r *http.Request) {
	// 这里需要说明的是发来的数据是通过post发过来一个json格式的数据
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("read all failded, ", err)
		responseError(w, 1001)
		return
	}

	var req model.Long2ShortRequest
	// 将反序列化的数据保存在结构体中
	err = json.Unmarshal(data, &req)
	if err != nil {
		fmt.Println("Unmarshal failded, ", err)
		responseError(w, 1002)
		return
	}

	resp, err := logic.Long2Short(&req)
	if err != nil {
		fmt.Println("Long2Short failded, ", err)
		responseError(w, 1003)
		return
	}

	responseSuccess(w, resp)
}

// 短地址到长地址
func Short2Long(w http.ResponseWriter, r *http.Request) {
	// 这里需要说明的是发来的数据是通过post发过来一个json格式的数据
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("read all failded, ", err)
		responseError(w, 1001)
		return
	}

	var req model.Short2LongRequest
	// 将反序列化的数据保存在结构体中
	err = json.Unmarshal(data, &req)
	if err != nil {
		fmt.Println("Unmarshal failded, ", err)
		responseError(w, 1002)
		return
	}

	resp, err := logic.Short2Long(&req)
	if err != nil {
		fmt.Println("Long2Short failded, ", err)
		responseError(w, 1003)
		return
	}
	responseSuccess(w, resp)
}

func main(){
	err := logic.InitDb("root:123456@tcp(192.168.50.145:3306)/short_url?parseTime=true")
	if err != nil{
		fmt.Printf("init db failed,err:%v\n",err)
		return
	}
	http.HandleFunc("/trans/long2short", Long2Short)
	http.HandleFunc("/trans/short2long", Short2Long)
	http.ListenAndServe(":18888", nil)
}
