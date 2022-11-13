package engine

import (
	"fmt"
	"github.com/dreamerjackson/newbiecrawler/collect"
	"github.com/dreamerjackson/newbiecrawler/parse"
)

type Crawler struct {
	Fetcher collect.Fetcher
}

func (c *Crawler) Start(seeds []string) {
	for _, url := range seeds {
		body, err := c.Fetcher.Get(url)
		if err != nil {
			fmt.Println("get content failed:%v", err)
		}
		if len(body) < 4096 {
			fmt.Println("get content failed: length body < 4096")
		} else {
			parse.ParseContent(body)
		}
	}
}
