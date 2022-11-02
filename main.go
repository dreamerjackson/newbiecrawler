package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	urls := []string{"https://www.thepaper.cn/", "https://www.douban.com/group/szsh"}

	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("fetch url error:%v\n", err)
			return
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Error status code:%v\n", resp.StatusCode)
		}

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println("read content failed:%v", err)
			return
		}

		if len(body) < 4096 {
			fmt.Println("get content failed:%v", "length body < 4096")
		} else {
			fmt.Println("body:", len(body))
		}
	}
}
