package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36"),
		//colly.Debugger(&debug.LogDebugger{}),
	)

	c.AllowedDomains = []string{"exhentai.org"}

	c.OnHTML(".gl3t a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println("Link: ", e.Text, link)
	})
	
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Host", "exhentai.org")
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
		r.Headers.Set("Accept-Encoding", "gzip, deflate, br")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,ja;q=0.7")
		r.Headers.Set("Cache-Control", "max-age=0")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Cookie", "ipb_member_id=1044069; ipb_pass_hash=4268d5c05c94cdedf124eabf2f1e7c95; igneous=45e5748db; sk=3hl0ggzrgfvcsp3wdu4tarft1k7v")
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
	})

	c.Visit("https://exhentai.org/")

}