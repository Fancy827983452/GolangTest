package util

import "encoding/json"

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
