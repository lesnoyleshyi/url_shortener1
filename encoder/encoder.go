package encoder

import (
	"math/rand"
	"time"
)

type Url struct {
	UrlShort string `json:"url_short"`
	UrlLong  string `json:"url_long"`
}

func (u *Url) Encode() {
	var alphabet = []rune("0123456789_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var shortUrlLength = 10
	rand.Seed(time.Now().UnixNano())

	encodedArr := make([]rune, shortUrlLength)
	for i := range encodedArr {
		encodedArr[i] = alphabet[rand.Intn(len(alphabet))]
	}
	u.UrlShort = string(encodedArr)
}
