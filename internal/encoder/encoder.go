package encoder

import (
	"crypto/md5"

	"github.com/ZaX51/url-shortener/internal/base62"
)

func Encode(text string, length int) string {
	crc := md5.Sum([]byte(text))

	return base62.ToBase62(crc[:])[:length]
}
