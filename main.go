package main

import (
	"fmt"
	"github.com/dreamerjackson/newbiecrawler/collect"
	"regexp"
)

func main() {
	urls := []string{"https://book.douban.com"}
	for _, url := range urls {
		body, err := collect.GetByBrowserFetch(url)
		if err != nil {
			fmt.Println("get content failed:%v", err)
		}
		if len(body) < 4096 {
			fmt.Println("get content failed: length body < 4096")
		} else {
			parseContent(body)
		}
	}
}

func parseContent(content []byte) {
	//<a href="/tag/科普" class="tag">科普</a>
	re := regexp.MustCompile(`<a href="([^"]+)" class="tag">[^<]+</a>`)

	matches := re.FindAllSubmatch(content, -1)

	for _, m := range matches {
		fmt.Printf("url:%s\n", "https://book.douban.com"+string(m[1]))
	}
}
