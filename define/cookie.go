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



	return cookies
}