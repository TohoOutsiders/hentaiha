package util

import (
	"github.com/gocolly/colly"
	"os"
	"sync"
)

type Tools struct {
}

var (
	Tool = New()
	once sync.Once
)

func New() (t *Tools) {
	once.Do(func() {
		t = &Tools{}
	})
	return t
}

func (t *Tools) CheckDirExist(path string) {
	_, err := os.Stat(path)

	if !os.IsNotExist(err) {
		os.RemoveAll(path)
	}
	os.MkdirAll(path, os.ModePerm)
}

func (t *Tools) SetHeader(r *colly.Request) {
	r.Headers.Set("Host", "exhentai.org")
	r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	r.Headers.Set("Accept-Encoding", "gzip, deflate, br")
	r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,ja;q=0.7")
	r.Headers.Set("Cache-Control", "max-age=0")
	r.Headers.Set("Connection", "keep-alive")
	r.Headers.Set("Cookie", "ipb_member_id=1044069; ipb_pass_hash=4268d5c05c94cdedf124eabf2f1e7c95; igneous=45e5748db; sk=3hl0ggzrgfvcsp3wdu4tarft1k7v")
}
