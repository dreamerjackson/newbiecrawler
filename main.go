package main

import (
	"fmt"
	"github.com/dreamerjackson/newbiecrawler/collect"
	"github.com/dreamerjackson/newbiecrawler/engine"
	"github.com/dreamerjackson/newbiecrawler/parse"
	"github.com/dreamerjackson/newbiecrawler/proxy"
	"time"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	before := time.Now()
	defer func() {
		fmt.Println(time.Since(before).Milliseconds())
	}()

	proxyURLs := []string{"http://127.0.0.1:8888"}
	p, err := proxy.RoundRobinProxySwitcher(proxyURLs...)
	if err != nil {
		fmt.Println("RoundRobinProxySwitcher failed")
	}

	f := &collect.BrowserFetch{
		Proxy: p,
	}
	c := engine.NewCrawler(f)
	c.Start([]*collect.Request{
		{
			Url:       "https://book.douban.com",
			Cookie:    "",
			ParseFunc: parse.ParseURL,
		},
	})
}
