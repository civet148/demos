package main

import (
	"github.com/civet148/gotools/log"
	"github.com/gocolly/colly/v2"
)

func main() {
	var err error
	c := colly.NewCollector()
	if err = c.Post("https://mail.126.com/", map[string]string{"username": "civet148@126.com", "password": "li@bin148"}); err != nil {
		log.Errorf("%s", err)
		return
	}
}
