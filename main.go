package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"hentai/util"
	"os"
	"strconv"
	"strings"
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

	//c.AllowedDomains = []string{"exhentai.org"}

	cheeioListPage(c)
}

/**
	请求根地址，列表页
 */
func cheeioListPage(c *colly.Collector) {
	c.OnRequest(func(r *colly.Request) {
		util.New().SetHeader(r)
		fmt.Println("Visiting", r.URL.String())
	})

	// 获取每个list .gl3t a标签元素的详情链接
	c.OnHTML(".gl3t a[href]", func(e *colly.HTMLElement) {
		d := c.Clone()
		requestDetailPage(d, e)
	})

	c.OnResponse(func(r *colly.Response) {
	})

	c.Visit("https://exhentai.org/?f_search=chinese")
}

/**
	列表各个详情页访问
 */
func requestDetailPage(c *colly.Collector, cheeioEl *colly.HTMLElement)  {
	var err error
	var title *string
	link := cheeioEl.Attr("href")

	c.OnRequest(func(dr *colly.Request) {
		util.New().SetHeader(dr)
		fmt.Println("Detail Visiting", dr.URL.String())
	})

	// 获取标题名称创建漫画文件夹
	c.OnHTML("#gn", func(de *colly.HTMLElement) {
		title = &de.Text
		groupDir := fmt.Sprintf("%s%s", outputDir, *title)
		err = os.Mkdir(groupDir, os.ModePerm)
		if err != nil {
			fmt.Errorf("%v\n", err)
		}
		//de.Response.Save(groupDir)
		fmt.Println("detail title: ", de.Text)

	})

	// 获取详情第一页图片链接
	c.OnHTML(".gdtl:first-child > a[href]", func(dde *colly.HTMLElement) {
		d := c.Clone()
		mapImageForHentai(d, dde, title)
	})

	fmt.Printf("End Request\n\n")
	c.Visit(cheeioEl.Request.AbsoluteURL(link))
}

func mapImageForHentai(c *colly.Collector, detailEl *colly.HTMLElement, title *string) {
	var num *int
	index := 2
	link := detailEl.Attr("href")

	c.OnRequest(func(mr *colly.Request) {
		util.New().SetHeader(mr)
		fmt.Println("Picture Visiting", mr.URL.String())
	})

	c.OnHTML("#i2 .sn span:last-child", func(e *colly.HTMLElement) {
		atoiNum, _ := strconv.Atoi(e.Text)
		num = &atoiNum
	})

	// 获取图片存储图片
	c.OnHTML("#i3 > a > img", func(e *colly.HTMLElement) {
		imgSrc := e.Attr("src")
		fmt.Println("Image Src: ", imgSrc)
		d := c.Clone()
		reduceImage(d, e, title)
	})
	
	c.OnHTML("#i3 > a[href]", func(e *colly.HTMLElement) {
		next := e.Attr("href")
		fmt.Println("Next Image Page: ", next)
		d := c.Clone()
		nextImageGiveMe(d, e, title, &index, *num)
	})

	c.Visit(detailEl.Request.AbsoluteURL(link))
}

func nextImageGiveMe(c *colly.Collector, mapEl *colly.HTMLElement, title *string, nextIndex *int, total int) {
	fmt.Println("=======================index ", *nextIndex, "=========================")
	fmt.Println("=======================total ", total, "=========================")
	if *nextIndex <= total {
		next := mapEl.Attr("href")
		fmt.Println("Next Image Page: ", next)

		c.OnRequest(func(mr *colly.Request) {
			util.New().SetHeader(mr)
			fmt.Println("Picture Visiting", mr.URL.String())
		})

		// 获取图片存储图片
		c.OnHTML("#i3 > a > img", func(e *colly.HTMLElement) {
			imgSrc := e.Attr("src")
			fmt.Println("Image Src: ", imgSrc)
			d := c.Clone()
			reduceImage(d, e, title)
		})

		c.OnHTML("#i3 > a[href]", func(e *colly.HTMLElement) {
			next := e.Attr("href")
			fmt.Println("Next Image Page: ", next)
			d := c.Clone()
			*nextIndex++
			nextImageGiveMe(d, mapEl, title, nextIndex, total)
		})

		c.Visit(next)
	}
}

func reduceImage(c *colly.Collector, mapEl *colly.HTMLElement, title *string) {
	link := mapEl.Attr("src")

	c.OnRequest(func(rr *colly.Request) {
		util.New().SetHeader(rr)
		fmt.Println("Download Visiting", rr.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Errorf("it's me!!\n")
		if strings.Index(r.Headers.Get("Content-Type"), "image") > -1 {
			r.Save(fmt.Sprintf("%s%s/", outputDir, *title) + r.FileName())
			return
		}
	})

	c.Visit(link)
}
