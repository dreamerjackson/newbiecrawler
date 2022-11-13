package parse

import (
	"fmt"
	"regexp"
	"strconv"
)

var autoRe = regexp.MustCompile(`<span class="pl"> 作者</span>:[\d\D]*?<a.*?>([^<]+)</a>`)
var public = regexp.MustCompile(`<span class="pl">出版社:</span>([^<]+)<br/>`)
var pageRe = regexp.MustCompile(`<span class="pl">页数:</span> ([^<]+)<br/>`)
var priceRe = regexp.MustCompile(`<span class="pl">定价:</span>([^<]+)<br/>`)
var scoreRe = regexp.MustCompile(`<strong class="ll rating_num " property="v:average">([^<]+)</strong>`)
var intoRe = regexp.MustCompile(`<div class="intro">[\d\D]*?<p>([^<]+)</p></div>`)

type Bookdetail struct {
	BookName  string
	Author    string
	Publicer  string
	Bookpages int
	Price     string
	Score     string
	Into      string
}

func (b Bookdetail) String() string {
	return "书籍名字:" + b.BookName + " 作者 :" + b.Author + " 出版社" + b.Publicer + " 书籍页数：" + strconv.Itoa(b.Bookpages) + " 价格：" + b.Price + " 得分" + b.Score + " \n简介:" + b.Into
}

func ParseContent(content []byte) {
	bookdetail := Bookdetail{}
	bookdetail.Author = ExtraString(content, autoRe)
	page, err := strconv.Atoi(ExtraString(content, pageRe))

	if err == nil {
		bookdetail.Bookpages = page
	}
	bookdetail.Publicer = ExtraString(content, public)
	bookdetail.Into = ExtraString(content, intoRe)
	bookdetail.Score = ExtraString(content, scoreRe)
	bookdetail.Price = ExtraString(content, priceRe)

	fmt.Println(bookdetail)
}

func ExtraString(contents []byte, re *regexp.Regexp) string {

	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}

}
