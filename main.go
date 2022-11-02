package main

import (
	"fmt"
	"github.com/dreamerjackson/newbiecrawler/collect"
)

func main() {
	urls := []string{"https://www.thepaper.cn/", "https://www.douban.com/group/szsh"}
	for _, url := range urls {
		body, err := collect.Get(url)
		if err != nil {
			fmt.Println("get content failed:%v", err)
		}
		if len(body) < 4096 {
			fmt.Println("get content failed: length body < 4096")
		} else {
			fmt.Println("body:", len(body))
		}
	}
}
