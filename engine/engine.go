package engine

import (
	"fmt"
	"github.com/dreamerjackson/newbiecrawler/collect"
)

type Crawler struct {
	out       chan collect.ParseResult
	Fetcher   collect.Fetcher
	scheduler Scheduler
}

func NewCrawler(f collect.Fetcher) *Crawler {
	c := &Crawler{}

	c.Fetcher = f
	c.scheduler = NewSchedule()
	c.out = make(chan collect.ParseResult)
	return c
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

func (c *Crawler) CreateWork() {

	for {
		req := c.scheduler.Pull()
		body, err := c.Fetcher.Get(req.Url)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if len(body) < 6000 {
			fmt.Println("can't fetch")
			continue
		}
		result := req.ParseFunc(body, req)

		if len(result.Requesrts) > 0 {
			go c.scheduler.Push(result.Requesrts...)
		}

		c.out <- result
	}
}

func (c *Crawler) HandleResult() {
	for result := range c.out {
		for _, item := range result.Items {
			fmt.Println(item)
		}
	}
}
