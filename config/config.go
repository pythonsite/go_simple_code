package config

import (
	"sync"
	"os"
	"bufio"
	"io"
	"strings"
	"fmt"
	"strconv"
	"time"
)

type Config struct {
	filename string
	lastModifyTime int64
	data map[string]string
	rwLock sync.RWMutex
	notifyList []Notifyer
}

func NewConfig(filename string)(conf *Config,err error){
	conf = &Config{
		filename:filename,
		data:make(map[string]string,1024),
	}
	m,err := conf.parse()
	if err != nil{
		return
	}
	conf.rwLock.Lock()
	conf.data = m
	conf.rwLock.Unlock()
	go conf.reload()
	return
}

func (c *Config) AddNotifyer(n Notifyer){
	c.notifyList = append(c.notifyList,n)
}

func (c *Config) reload(){
	// 这里启动一个定时器，每5秒重新加载一次配置文件
	ticker := time.NewTicker(time.Second*5)
	for _ = range ticker.C{
		func(){
			file,err := os.Open(c.filename)
			if err != nil{
				fmt.Printf("open %s failed,err:%v\n",c.filename,err)
				return
			}
			defer file.Close()
			fileInfo,err := file.Stat()
			if err != nil{
				fmt.Printf("stat %s failed,err:%v\n",c.filename,err)
				return
			}
			curModifyTime := fileInfo.ModTime().Unix()
			fmt.Printf("%v --- %v\n",curModifyTime,c.lastModifyTime)
			if curModifyTime > c.lastModifyTime{
				m,err := c.parse()
				if err != nil{
					fmt.Println("parse failed,err:",err)
					return
				}
				c.rwLock.Lock()
				c.data = m
				c.rwLock.Unlock()
				for _, n:=range c.notifyList{
					n.Callback(c)
				}
				c.lastModifyTime = curModifyTime
			}
		}()
	}
}

func (c *Config) parse()(m map[string]string,err error){
	// 读文件并或将文件中的数据以k/v的形式存储到map中
	m = make(map[string]string,1024)
	file,err := os.Open(c.filename)
	if err != nil{
		return
	}
	var lineNo int
	reader := bufio.NewReader(file)
	for{
		// 一行行的读文件
		line,errRet := reader.ReadString('\n')
		if errRet == io.EOF{
			// 表示读到文件的末尾
			break
		}
		if errRet != nil{
			// 表示读文件出问题
			err = errRet
			return
		}
		lineNo++
		line = strings.TrimSpace(line) // 取出空格
		if len(line) == 0 || line[0] == '\n' || line[0] == '+' || line[0] == ';'{
			// 当前行为空行或者是注释行等
			continue
		}
		arr := strings.Split(line,"=") // 通过=进行切割取出k/v结构
		if len(arr) == 0{
			fmt.Printf("invalid config,line:%d\n",lineNo)
			continue
		}
		key := strings.TrimSpace(arr[0])
		if len(key) == 0{
			fmt.Printf("invalid config,line:%d\n",lineNo)
			continue
		}
		if len(arr) == 1{
			m[key] = ""
			continue
		}
		value := strings.TrimSpace(arr[1])
		m[key] = value
	}
	return
}

func (c *Config) GetInt(key string)(value int,err error){
	// 根据int获取
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	str,ok:=c.data[key]
	if !ok{
		err = fmt.Errorf("key[%s] not found",key)
		return
	}
	value,err = strconv.Atoi(str)
	return
}

func (c *Config)GetIntDefault(key string,defval int)(value int){
	// 默认值
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	str,ok:=c.data[key]
	if !ok{
		value = defval
		return
	}
	value,err := strconv.Atoi(str)
	if err != nil{
		value =defval
		return
	}
	return
}

func (c *Config) GetString(key string)(value string,err error){
	// 根据字符串获取
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	value,ok := c.data[key]
	if !ok{
		err = fmt.Errorf("key[%s] not found",key)
		return
	}
	return
}
