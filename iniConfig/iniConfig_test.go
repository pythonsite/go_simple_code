package iniConfig

import (
	"fmt"
	"io/ioutil"
	"testing"
)

type ServerConfig struct {
	Ip string	`ini:"ip"`
	Port int	`ini:"port"`
}

type MysqlConfig struct {
	UserName string	`ini:"username"`
	Passwd string	`ini:"passwd"`
	DataBase string `ini:"database"`
	Host string	`ini:"host"`
	Port int	`ini:"port"`
	Timeout float32 `ini:"timeout"`
}

type Config struct {
	ServerConf ServerConfig `ini:"server"`
	MysqlConf MysqlConfig	`ini:"mysql"`
}

func TestIniConfig(t *testing.T) {
	fmt.Println("hello")
	data, err := ioutil.ReadFile("./config.ini")
	if err != nil {
		t.Error("open file error:",err)
	}
	var conf Config
	err = unmarshal(data, &conf)
	if err != nil {
		t.Fatalf("unmarshal failed, err:%v",err)
	}
	t.Log("unmarshall success")
	t.Logf("unmarshal success, conf:%#v\n",conf)

	confData , err := marshal(conf)
	if err != nil {
		t.Fatalf("marshal error:%v",err)
	}
	t.Logf("marshal success: confData:\n%v", string(confData))

	err = MarshalFile("./test.ini", conf)
	if err != nil {
		t.Fatalf("MarshalFile error:%v",err)
	}
}

func TestUnmarshalFile(t *testing.T) {
	serverConfig := &ServerConfig{
		Ip: "192.168.11.1",
		Port: 8080,
	}
	mysqlConfig := &MysqlConfig{
		UserName: "zhaofan",
		Passwd:"root123",
		DataBase:"mydb",
		Host:"127.0.0.1",
		Port:3309,
		Timeout:1.222,
	}
	conf := Config{
		ServerConf:*serverConfig,
		MysqlConf:*mysqlConfig,
	}
	err := MarshalFile("./test.ini", conf)
	if err != nil {
		t.Errorf("marshalfile failed,err:%v",err)
		return
	}
	var conf2 Config
	err =UnmarshalFile("./test.ini", &conf2)
	if err != nil {
		t.Errorf("unmarshalfile failed,err:%v",err)
		return
	}
	t.Logf("unmarhsal success, conf:%v", conf2)
}