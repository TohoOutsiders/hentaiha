package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"hentai/util"
	"os"
	"strconv"
	"strings"
)

var logger util.ILogger

const (
	outputDir = "./out/"
	pageTotal = 5
)

func init()  {
	logger = &util.Logger{}
}

func main() {
	util.New().CheckDirExist(outputDir)

	chNum := pageTotal + 1
	ch := make(chan int, chNum)

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36"),
		//colly.Debugger(&debug.LogDebugger{}),
	)

	go cheerioFirstListPage(c, ch)

	i := 1
	for {
		if i > 5 {
			break
		}
		go cheerioListPage(c, i, ch)
		i++
	}

	cashPool :=0
	for cashPool < chNum {
		cashPool += <-ch
	}
}

/**
	请求根地址，列表页
 */
func cheerioFirstListPage(c *colly.Collector, ch chan int) {
	c.OnRequest(func(r *colly.Request) {
		util.New().SetHeader(r)
		logger.Info("Visiting", r.URL.String())
	})

	// 获取每个list .gl3t a标签元素的详情链接
	c.OnHTML(".gl3t a[href]", func(e *colly.HTMLElement) {
		d := c.Clone()
		requestDetailPage(d, e)
	})

	c.OnResponse(func(r *colly.Response) {
	})

	c.Visit("https://exhentai.org/?f_search=chinese")

	ch <- 1
}

func cheerioListPage(c *colly.Collector, pageIndex int, ch chan int) {
	c.OnRequest(func(r *colly.Request) {
		util.New().SetHeader(r)
		logger.Info("Visiting", r.URL.String())
	})

	// 获取每个list .gl3t a标签元素的详情链接
	c.OnHTML(".gl3t a[href]", func(e *colly.HTMLElement) {
		d := c.Clone()
		requestDetailPage(d, e)
	})

	c.OnResponse(func(r *colly.Response) {
	})

	c.Visit(fmt.Sprintf("https://exhentai.org/?page=%d&f_search=chinese", pageIndex))

	ch <- 1
}

/**
	列表各个详情页访问
 */
func requestDetailPage(c *colly.Collector, cheeioEl *colly.HTMLElement)  {
	var title *string
	link := cheeioEl.Attr("href")

	c.OnRequest(func(dr *colly.Request) {
		util.New().SetHeader(dr)
		logger.Normal("Detail Visiting", dr.URL.String())
	})

	// 获取标题名称创建漫画文件夹
	c.OnHTML("#gn", func(de *colly.HTMLElement) {
		title = &de.Text
		groupDir := fmt.Sprintf("%s%s", outputDir, *title)
		_ = os.Mkdir(groupDir, os.ModePerm)
		logger.Normal("detail title: ", de.Text)
	})

	// 获取详情第一页图片链接
	c.OnHTML(".gdtl:first-child > a[href]", func(dde *colly.HTMLElement) {
		d := c.Clone()
		mapImageForHentai(d, dde, title)
	})

	logger.Complate("End Request")
	c.Visit(cheeioEl.Request.AbsoluteURL(link))
}

func mapImageForHentai(c *colly.Collector, detailEl *colly.HTMLElement, title *string) {
	var num *int
	index := 2
	link := detailEl.Attr("href")

	c.OnRequest(func(mr *colly.Request) {
		util.New().SetHeader(mr)
	})

	c.OnHTML("#i2 .sn span:last-child", func(e *colly.HTMLElement) {
		atoiNum, _ := strconv.Atoi(e.Text)
		num = &atoiNum
	})

	// 获取图片存储图片
	c.OnHTML("#i3 > a > img", func(e *colly.HTMLElement) {
		d := c.Clone()
		reduceImage(d, e, title)
	})
	
	c.OnHTML("#i3 > a[href]", func(e *colly.HTMLElement) {
		next := e.Attr("href")
		logger.Info("【HEAD】Start Image Download！！", next)
		d := c.Clone()
		nextImageGiveMe(d, e, title, &index, *num)
	})

	c.Visit(detailEl.Request.AbsoluteURL(link))
}

func nextImageGiveMe(c *colly.Collector, mapEl *colly.HTMLElement, title *string, nextIndex *int, total int) {
	if *nextIndex == total {
		logger.Underline("=======================>  total ", total, )
	}
	if *nextIndex <= total {
		next := mapEl.Attr("href")
		c.OnRequest(func(mr *colly.Request) {
			util.New().SetHeader(mr)
		})

		// 获取图片存储图片
		c.OnHTML("#i3 > a > img", func(e *colly.HTMLElement) {
			d := c.Clone()
			reduceImage(d, e, title)
		})

		c.OnHTML("#i3 > a[href]", func(e *colly.HTMLElement) {
			next := e.Attr("href")
			logger.Info("Next Image Page: ", next)
			d := c.Clone()
			*nextIndex++
			nextImageGiveMe(d, e, title, nextIndex, total)
		})

		c.Visit(mapEl.Request.AbsoluteURL(next))
	}
}

func reduceImage(c *colly.Collector, mapEl *colly.HTMLElement, title *string) {
	link := mapEl.Attr("src")

	c.OnRequest(func(rr *colly.Request) {
		util.New().SetHeader(rr)
	})

	c.OnResponse(func(r *colly.Response) {
		if strings.Index(r.Headers.Get("Content-Type"), "image") > -1 {
			r.Save(fmt.Sprintf("%s%s/", outputDir, *title) + r.FileName())
			return
		}
	})

	c.Visit(link)
}
