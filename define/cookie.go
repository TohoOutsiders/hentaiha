package define

import (
	"net/http"
	"time"
)

const (
	MAX_AGE = time.Hour * 24 / time.Second
	maxAge = int(MAX_AGE)
)

func GetCookies() []*http.Cookie {
	cookies := make([]*http.Cookie, 4)

	cookies[0] = &http.Cookie{
		Name: "igneous",
		Value: "45e5748db",
		Path: "/",
		Domain: ".exhentai.org",
		MaxAge: maxAge,
	}

	cookies[1] = &http.Cookie{
		Name: "ipb_member_id",
		Value: "1044069",
		Path: "/",
		Domain: ".exhentai.org",
		MaxAge: maxAge,
	}

	cookies[2] = &http.Cookie{
		Name: "ipb_pass_hash",
		Value: "4268d5c05c94cdedf124eabf2f1e7c95",
		Path: "/",
		Domain: ".exhentai.org",
		MaxAge: maxAge,
	}

	cookies[3] = &http.Cookie{
		Name: "sk",
		Value: "3hl0ggzrgfvcsp3wdu4tarft1k7v",
		Path: "/",
		Domain: ".exhentai.org",
		MaxAge: maxAge,
	}

	return cookies
}