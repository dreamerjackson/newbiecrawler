package parse

import (
	"github.com/dreamerjackson/newbiecrawler/collect"
	"github.com/dreamerjackson/newbiecrawler/db"
	"regexp"
	"strconv"
)

const urlListRe = `<a href="([^"]+)" class="tag">([^<]+)</a>`

func ParseURL(contents []byte, url *collect.Request) collect.ParseResult {
	re := regexp.MustCompile(urlListRe)

	matches := re.FindAllSubmatch(contents, -1)
	result := collect.ParseResult{}

	for _, m := range matches[:1] {
		u := "https://book.douban.com" + string(m[1])
		result.Requesrts = append(
			result.Requesrts, &collect.Request{
				Url: u,
				ParseFunc: func(c []byte, request *collect.Request) collect.ParseResult {
					return ParseBookList(c, nil)
				},
			})
	}
	return result
}

const BooklistRe = `<a.*?href="([^"]+)" title="([^"]+)"`

func ParseBookList(contents []byte, url *collect.Request) collect.ParseResult {
	re := regexp.MustCompile(BooklistRe)

	matches := re.FindAllSubmatch(contents, -1)
	result := collect.ParseResult{}

	for _, m := range matches[:2] {
		u := string(m[1])
		result.Requesrts = append(
			result.Requesrts, &collect.Request{
				Url: u,
				ParseFunc: func(c []byte, request *collect.Request) collect.ParseResult {
					return GetContent(c, &collect.Request{
						Temp: string(m[2]),
					})
				},
			})
	}
	return result
}

var autoRe = regexp.MustCompile(`<span class="pl"> 作者</span>:[\d\D]*?<a.*?>([^<]+)</a>`)
var public = regexp.MustCompile(`<span class="pl">出版社:</span>([^<]+)<br/>`)
var pageRe = regexp.MustCompile(`<span class="pl">页数:</span> ([^<]+)<br/>`)
var priceRe = regexp.MustCompile(`<span class="pl">定价:</span>([^<]+)<br/>`)
var scoreRe = regexp.MustCompile(`<strong class="ll rating_num " property="v:average">([^<]+)</strong>`)
var intoRe = regexp.MustCompile(`<div class="intro">[\d\D]*?<p>([^<]+)</p></div>`)

func GetContent(content []byte, u *collect.Request) collect.ParseResult {
	bookdetail := db.Bookdetail{}

	bookdetail.BookName = u.Temp

	bookdetail.Author = ExtraString(content, autoRe)
	page, err := strconv.Atoi(ExtraString(content, pageRe))

	if err == nil {
		bookdetail.Bookpages = page
	}
	bookdetail.Publicer = ExtraString(content, public)
	bookdetail.Into = ExtraString(content, intoRe)
	bookdetail.Score = ExtraString(content, scoreRe)
	bookdetail.Price = ExtraString(content, priceRe)

	result := collect.ParseResult{
		Items: []interface{}{bookdetail},
	}

	return result
}

func ExtraString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
