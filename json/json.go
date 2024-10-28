package json

import (
	"github.com/bytedance/sonic"
	"github.com/cloudwego/gjson"
)

func Marshal(v interface{}) ([]byte, error) {
	return sonic.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return sonic.Unmarshal(data, v)
}

func UnmarshalString(data string, v interface{}) error {
	return sonic.UnmarshalString(data, v)
}

func Get(json string, path string) gjson.Result {
	return gjson.Get(json, path)
}
