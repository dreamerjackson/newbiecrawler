package engine

import (
	"fmt"
	"github.com/dreamerjackson/newbiecrawler/collect"
)

type Crawler struct {
	Fetcher collect.Fetcher
}

func (c *Crawler) Start(reqs []*collect.Request) {
	for i := 0; i < len(reqs); i++ {
		r := reqs[i]
		body, err := c.Fetcher.Get(r.Url)
		if err != nil {
			fmt.Println("get content failed:%v", err)
		}
		if len(body) < 4096 {
			fmt.Println("get content failed: length body < 4096")
		} else {
			results := r.ParseFunc(body, r)
			if results.Requesrts != nil {
				reqs = append(reqs, results.Requesrts...)
			}
			for _, item := range results.Items {
				fmt.Println(item)
			}
		}
	}
}
