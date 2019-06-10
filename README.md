# 変態は！！

### You need to do（废弃，使用公共用户Cookie，您也可以使用个人私有的Cookie）

```shell
# dir: project/util/tools.go
func (t *Tools) SetHeader(r *colly.Request) {
	. . .
	
	r.Headers.Set("Cookie", "add your cookie")
}
```

### $
```shell
go mod download

go run main.go

go build main.go
```

### EndTime

- 请求次数过多，导致封禁IP 40min(?)，可能更久或永封

### TODO

- [x] 搜索功能
- [ ] 使用配置文件(eg: 爬取的页数 ...)
- [ ] <del>登陆E-Hentai Forums保存并获取cookie</del>(使用公共用户Cookie)
- [ ] 终端输出美化

### contact me

- QQ Group: 795711415

- E-mail: gutrse3321@live.com
