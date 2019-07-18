package util

import (
	"github.com/gocolly/colly"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

type Tools struct {
}

var (
	Tool   = New()
	once   sync.Once
	logger ILogger
)

func init() {
	logger = &Logger{}
}

func New() (t *Tools) {
	once.Do(func() {
		t = &Tools{}
	})
	return t
}

func (t *Tools) ReplaceAll(s, old, newS string) (result string) {
	result = strings.Replace(s, old, newS, -1)
	return
}

func (t *Tools) CheckDirExist(path string) (bool, error) {
	logger.Normal("START FILE DIR CHECK ... Please waiting for me ...")
	_, err := os.Stat(path)

	if !os.IsNotExist(err) {
		dirList, e := ioutil.ReadDir(path)
		if e != nil {
			log.Panic(e)
		}
		for _, v := range dirList {
			os.RemoveAll(path + v.Name())
		}
	} else {
		os.Mkdir(path, os.ModePerm)
	}
	logger.Normal("End FILE DIR CHECK!!!")
	return true, nil
}

func (t *Tools) SetHeader(r *colly.Request) {
	r.Headers.Set("Host", "exhentai.org")
	r.Headers.Set("Referer", "https://exhentai.org/")
	r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	r.Headers.Set("Accept-Encoding", "gzip, deflate, br")
	r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,ja;q=0.7")
	r.Headers.Set("Cache-Control", "max-age=0")
	r.Headers.Set("Connection", "keep-alive")
	r.Headers.Set("Upgrade-Insecure-Requests", "1")
	r.Headers.Set("Cookie", "ipb_member_id=4483572; ipb_pass_hash=b1d7d5acd649a01a1643124c8a0918a8; igneous=df9724040; sk=3hl0ggzrgfvcsp3wdu4tarft1k7v")
}

func (t *Tools) ReadyGo(s int) {
	ch := t.ticker(s)
	time.Sleep(time.Duration(s) * time.Second)
	ch <- true
	close(ch)
}

func (t *Tools) ticker(s int) chan bool {
	ticker := time.NewTicker(time.Second)
	stopChan := make(chan bool)
	go func(ticker *time.Ticker) {
		num := s
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				logger.Underline(num)
				num--
			case stop := <-stopChan:
				if stop {
					logger.Normal("========= Game Start =========")
					return
				}
			}
		}
	}(ticker)
	return stopChan
}
