package main

import (
	"github.com/civet148/gotools/log"
	jsoniter "github.com/json-iterator/go"
)

type JsonObject struct {
	Id   int32  `json:"id"`
	Name string `json:"string"`
}

func main() {

	var joEncode = JsonObject{
		Id:   10001,
		Name: "jo",
	}

	data, err := jsoniter.Marshal(&joEncode)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	log.Infof("encode [%s]", data)
	var joDecode JsonObject
	if err = jsoniter.Unmarshal(data, &joDecode); err != nil {
		log.Errorf(err.Error())
		return
	}
	log.Infof("decode1 [%+v]", joDecode)
}
