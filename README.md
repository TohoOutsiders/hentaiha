# 変態は！！

### You need to do

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

### TODO

- [ ] 使用配置文件(eg: 爬取的页数 ...)
- [ ] 登陆E-Hentai Forums保存并获取cookie
- [ ] 终端输出美化

### contact me

- QQ Group: 795711415

- E-mail: gutrse3321@live.com
