package json

import (
	jsoniter "github.com/json-iterator/go"
)

var jsonApi jsoniter.API

func init() {
	// 标准解析模式
	jsonApi = jsoniter.ConfigCompatibleWithStandardLibrary
}

func Unmarshal(j string, v interface{}) error {
	return jsonApi.UnmarshalFromString(j, v)
}

func Marshal(v interface{}) (string, error) {
	return jsonApi.MarshalToString(v)
}
