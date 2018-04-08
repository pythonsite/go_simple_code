package config

import (
	"strings"
	"errors"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"reflect"
	"fmt"
	"strconv"
)

type ConfigEngine struct {
	data map[interface{}]interface{}
}

//func NewConfigEngine() *ConfigEngine {
//	return &ConfigEngine{
//		data:make(map[interface{}]interface{},10),
//	}
//}

func (c *ConfigEngine) Load (path string) error {
	ext := c.guessFileType(path)
	if ext == "" {
		return errors.New("cant not load" + path + " config")
	}
	return c.loadFromYaml(path)
}

func (c *ConfigEngine) guessFileType(path string) string {
	s := strings.Split(path,".")
	ext := s[len(s) - 1]
	switch ext {
	case "yaml","yml":
		return "yaml"
	}
	return ""
}

func (c *ConfigEngine) loadFromYaml(path string) error {
	yamlS,readErr := ioutil.ReadFile(path)
	if readErr != nil {
		return readErr
	}
	err := yaml.Unmarshal(yamlS, &c.data)
	if err != nil {
		return errors.New("can not parse "+ path + " config" )
	}
	return nil
}


func (c *ConfigEngine) Get(name string) interface{}{
	path := strings.Split(name,".")
	data := c.data
	for key, value := range path {
		fmt.Println(data)
		v, ok := data[value]
		if !ok {
			break
		}
		if (key + 1) == len(path) {
			return v
		}
		if reflect.TypeOf(v).String() == "map[interface {}]interface {}"{
			data = v.(map[interface {}]interface {})
		}
	}
	return nil
}

func (c *ConfigEngine) GetString(name string) string {
	value := c.Get(name)
	switch value:=value.(type){
	case string:
		return value
	case bool,float64,int:
		return fmt.Sprint(value)
	default:
		return ""
	}
}

func (c *ConfigEngine) GetInt(name string) int {
	value := c.Get(name)
	switch value := value.(type){
	case string:
		i,_:= strconv.Atoi(value)
		return i
	case int:
		return value
	case bool:
		if value{
			return 1
		}
		return 0
	case float64:
		return int(value)
	default:
		return 0
	}
}

func (c *ConfigEngine) GetBool(name string) bool {
	value := c.Get(name)
	switch value := value.(type){
	case string:
		str,_:= strconv.ParseBool(value)
		return str
	case int:
		if value != 0 {
			return true
		}
		return false
	case bool:
		return value
	case float64:
		if value != 0.0 {
			return true
		}
		return false
	default:
		return false
	}
}

func (c *ConfigEngine) GetFloat64(name string) float64 {
	value := c.Get(name)
	switch value := value.(type){
	case string:
		str,_ := strconv.ParseFloat(value,64)
		return str
	case int:
		return float64(value)
	case bool:
		if value {
			return float64(1)
		}
		return float64(0)
	case float64:
		return value
	default:
		return float64(0)
	}
}

func (c *ConfigEngine) GetStruct(name string,s interface{}) interface{}{
	d := c.Get(name)
	switch d.(type){
	case string:
		c.setField(s,name,d)
	case map[interface{}]interface{}:
		//
	}
	return s
}

func (c *ConfigEngine) mapToStruct(m map[interface{}]interface{},s interface{}) interface{}{
	for key, value := range m {
		switch key.(type) {
		case string:
			c.setField(s,key.(string),value)
		}
	}
	return s
}

func (c *ConfigEngine) setField(obj interface{},name string,value interface{}) error {
	structValue := reflect.Indirect(reflect.ValueOf(obj))
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj",name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)

	if structFieldType.Kind() == reflect.Struct && val.Kind() == reflect.Map {
		vint := val.Interface()

		switch vint.(type) {
		case map[interface{}]interface{}:
			for key, value := range vint.(map[interface{}]interface{}) {
				c.setField(structFieldValue.Addr().Interface(), key.(string), value)
			}
		case map[string]interface{}:
			for key, value := range vint.(map[string]interface{}) {
				c.setField(structFieldValue.Addr().Interface(), key, value)
			}
		}

	} else {
		if structFieldType != val.Type() {
			return errors.New("Provided value type didn't match obj field type")
		}

		structFieldValue.Set(val)
	}

	return nil


}





