package engine

import (
	"fmt"
	"github.com/dreamerjackson/newbiecrawler/collect"
	"github.com/dreamerjackson/newbiecrawler/db"
)

type Crawler struct {
	out       chan collect.ParseResult
	Fetcher   collect.Fetcher
	scheduler Scheduler
	db        *db.MysqlDB
}

func NewCrawler(f collect.Fetcher) *Crawler {
	c := &Crawler{}
	c.Fetcher = f
	c.scheduler = NewSchedule()
	c.out = make(chan collect.ParseResult)
	d, err := db.OpenDB()

	err = d.InitTable(db.Bookdetail{})
	if err != nil {
		panic(err)
	}
	c.db = d
	if err != nil {
		panic(err)
	}
	return c
}

func (c *Crawler) Start(reqs []*collect.Request) {
	go c.scheduler.Schedule()
	c.scheduler.Push(reqs...)

	for i := 0; i < 5; i++ {
		go c.CreateWork()
	}
	c.HandleResult()
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
			book, ok := item.(db.Bookdetail)
			if ok {
				err := c.db.Insert(book)
				if err != nil {
					fmt.Println("Error inserting book:", err)
					continue
				}
				fmt.Println("Inserted book:", book)
			}
		}
	}
}
