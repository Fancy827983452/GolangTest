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

//func RecResultJsonToPlain(json_str string,key string) {
//	var recResult string
//	var dat map[string]interface{}
//	json.Unmarshal([]byte(json_str), &dat)
//	if v, ok := dat[key]; ok {
//		ws := v.([]interface{})
//		for wsItem := range ws {
//			wsMap := wsItem.(map[string]interface{})
//		}
//		recResult = key.(string)
//
//		for i,wsItem := range ws {
//			wsMap := wsItem.(map[string]interface{})
//			if vCw, ok := wsMap["cw"]; ok {
//				cw := vCw.([]interface{})
//				for i,cwItem := range cw {
//					cwItemMap := cwItem.(map[string]interface{})
//					if w, ok := cwItemMap["w"]; ok {
//						recResult = recResult + w.(string)
//					}
//				}
//			}
//		}
//	}
//	fmt.Println(recResult)
//}