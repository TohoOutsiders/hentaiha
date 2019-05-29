package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"hentai/util"
)

const (
	outputDir = "./out/"
)

func main() {
	util.New().CheckDirExist(outputDir)

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36"),
		//colly.Debugger(&debug.LogDebugger{}),
	)

	c.AllowedDomains = []string{"exhentai.org"}

	c.OnRequest(func(r *colly.Request) {
		util.New().SetHeader(r)
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnHTML(".gl3t a[href]", func(e *colly.HTMLElement) {
		d := c.Clone()
		link := e.Attr("href")
		d.OnRequest(func(dr *colly.Request) {
			util.New().SetHeader(dr)
			fmt.Println("Detail Visiting", dr.URL.String())
		})
		d.OnHTML("#gn", func(de *colly.HTMLElement) {
			fmt.Println("detail title: ", de.Text)
		})
		d.OnResponse(func(r *colly.Response) {
			r.Save(outputDir + r.FileName())
		})
		d.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnResponse(func(r *colly.Response) {
	})

	c.Visit("https://exhentai.org/?f_search=chinese")
}
