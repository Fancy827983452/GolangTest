package util

import (
	"encoding/json"
)

//将结构体转变为json格式
func ParseJson(v interface{}) (string) {
	jsonByte, err := json.Marshal(v)
	CheckErr(err)
	ch := make(chan string, 1)
	go func(c chan string, str string){
		c <- str
	}(ch, string(jsonByte))
	strData := <-ch
	return strData
}