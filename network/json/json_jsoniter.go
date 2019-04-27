package cljson

import (
	"fmt"
	"github.com/json-iterator/go"
)

func Marshal(data interface{}) ([]byte, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(data)
}

func Unmarshal(data []byte, v interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(data, v)
}

/**
@author xiaoma
@lastupdate 2019-03-10
序列化成字节返回
*/
func MarshalByte(data interface{}) []byte {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("序列化失败:%v \n", err)
		return nil
	}
	return b
}
