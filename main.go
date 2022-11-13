package main

import (
	"fmt"
	"github.com/dreamerjackson/newbiecrawler/collect"
	"github.com/dreamerjackson/newbiecrawler/engine"
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

	seeds := []string{"https://book.douban.com/subject/36104107/"}
	f := &collect.BrowserFetch{}
	c := engine.Crawler{f}
	c.Start(seeds)
}
