package main

import (
	"go_dev/13/config"
	"fmt"
	"time"
	"sync/atomic"
)

type AppConfig struct {
	port int
	nginxAddr string
}

type AppconfigMgr struct {
	config atomic.Value
}

var appConfigMgr = &AppconfigMgr{}


func(a *AppconfigMgr)Callback(conf *config.Config){

	var appConfig = &AppConfig{}

	port,err := conf.GetInt("server_port")
	if err != nil{
		fmt.Println("get port failed,err:",err)
		return
	}
	appConfig.port = port
	fmt.Println("port:",appConfig.port)
	nginxAddr,err := conf.GetString("nginx_addr")
	if err != nil{
		fmt.Println("get nginx addr failed,err:",err)
		return
	}
	appConfig.nginxAddr = nginxAddr
	fmt.Println("nginx addr :",appConfig.nginxAddr)

	appConfigMgr.config.Store(appConfig)

}

func run(){
	for {
		// 每5秒打印一次数据，查看自己更改配置文件后是否可以热刷新
		appConfig := appConfigMgr.config.Load().(*AppConfig)
		fmt.Println("port:",appConfig.port)
		fmt.Println("nginx addr:",appConfig.nginxAddr)
		time.Sleep(5* time.Second)
	}
}

func main() {
	conf,err := config.NewConfig("/Users/zhaofan/go_project/src/go_dev/13/config_test/config.conf")
	if err != nil{
		fmt.Println("parse config failed,err:",err)
		return
	}
	//打开文件获取内容后，将自己加入到被通知的切片中
	conf.AddNotifyer(appConfigMgr)

	var appConfig = &AppConfig{}

	appConfig.port,err = conf.GetInt("server_port")
	if err != nil{
		fmt.Println("get port failed,err:",err)
		return
	}
	fmt.Println("port:",appConfig.port)

	appConfig.nginxAddr,err = conf.GetString("nginx_addr")
	if err != nil{
		fmt.Println("get nginx addr failed,err:",err)
		return
	}
	fmt.Println("nginx addr:",appConfig.nginxAddr)
	appConfigMgr.config.Store(appConfig)
	run()

}