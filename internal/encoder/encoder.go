package encoder

import (
	"crypto/md5"
	"encoding/base32"
)

func Encode(s string) string {
	crc := md5.Sum([]byte(s))

	return base32.HexEncoding.EncodeToString(crc[:])
}
