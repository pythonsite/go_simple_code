package iniConfig

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func Marshal(data interface{})(result []byte,err error) {
	return
}

func Unmarshal(data []byte, result interface{}) (err error){
	typeInfo := reflect.TypeOf(result)
	if typeInfo.Kind() != reflect.Ptr {
		err = errors.New("please pass address")
		return
	}
	typeStruct := typeInfo.Elem()
	if typeStruct.Kind() != reflect.Struct {
		err = errors.New("please pass struct")
		return
	}
	lineArr := strings.Split(string(data),"\n")
	var lastFieldName string
	for index, line := range lineArr {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		// 如果是注释，忽略
		if line[0] == ';'|| line[0] == '#' {
			continue
		}
		if line[0] == '[' {
			lastFieldName, err = parseSection(line, typeStruct)
			if err != nil {
				err =fmt.Errorf("%v lineNo:%d",err,index+1)
				return
			}
			continue
		}
		err = parseItem(lastFieldName, line, result)
		if err != nil {
			err = fmt.Errorf("%v lineNo:%d",err, index + 1)
		}
	}
	return
}


func parseSection(line string, typeInfo2 reflect.Type) (fieldName string,err error) {

	if line[0] == '[' && len(line) <= 2 {
		err = fmt.Errorf("error invalid section:%s ", line)
		return
	}
	if line[0] == '[' && line[len(line) -1] != ']' {
		err = fmt.Errorf("error invalid section:%s", line)
		return
	}
	if line[0] == '[' && line[len(line) -1] == ']'{
		sectionName := strings.TrimSpace(line[1:len(line) -1])
		if len(sectionName) == 0 {
			err = fmt.Errorf("error invalid sectionName:%s", line,)
			return
		}
		for i := 0;i< typeInfo2.NumField();i ++ {
			field := typeInfo2.Field(i)
			tagValue := field.Tag.Get("ini")
			if tagValue == sectionName {
				fieldName = field.Name
				fmt.Println("field name:", fieldName)
				break
			}
		}
	}
	return
}

func parseItem(lastFieldName string, line string, result interface{}) (err error) {
	index := strings.Index(line, "=")
	if index == -1 {
		err = fmt.Errorf("sytax error, line:%s", line)
		return
	}
	return
}


