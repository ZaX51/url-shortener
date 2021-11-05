package encoder

import (
	"crypto/md5"

	"github.com/ZaX51/url-shortener/internal/base62"
)

func Encode(s string) string {
	crc := md5.Sum([]byte(s))

	return base62.ToBase62(crc[:])
}
