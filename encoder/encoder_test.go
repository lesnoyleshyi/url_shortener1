package encoder

import (
	"strings"
	"testing"
)

func TestUrl_Encode_len(t *testing.T) {
	var urlStruct Url
	urlStruct.UrlLong = "https://www.github.com/lesnoyleshyi"

	urlStruct.Encode()

	shortenedUrl := urlStruct.UrlShort
	//shortenedUrl := "so_small"

	if len(shortenedUrl) != 10 {
		t.Errorf("Wrong len of shortened url: %d, expected %d",
			len(shortenedUrl), 10)
	}
}

func TestUrl_Encode_allowed_symbols(t *testing.T) {
	var urlStruct = Url{UrlLong: "https://www.github.com/lesnoyleshyi"}
	alphabet := "0123456789_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	urlStruct.Encode()
	//urlStruct.UrlShort = "lol!"
	for _, symb := range urlStruct.UrlShort {
		if !strings.ContainsRune(alphabet, symb) {
			t.Errorf("Wrong symbol in shortened url: %c", symb)
		}
	}
}
