package main

import (
	"encoding/json"
	"fmt"
)

type JsonType struct {
	Key1 int    `json:"key1"`
	Key2 bool   `json:"key2"`
	Key3 string `json:"key3"`
}

func main() {
	sObj := JsonType{
		Key1: 1,
		Key2: true,
	}
	sObj.Key3 = "some value"
	jObj, _ := json.Marshal(&sObj)

	fmt.Println(string(jObj))
}
