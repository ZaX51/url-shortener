package base62

import (
	"math/big"
)

const (
	base         int64 = 62
	characterSet       = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func ToBase62(src []byte) string {
	result := ""
	bigIntBase := big.NewInt(base)
	remainder := new(big.Int)
	num := new(big.Int).SetBytes(src)
	zero := big.NewInt(0)

	for num.Cmp(zero) == +1 {
		remainder.Mod(num, bigIntBase)
		num.Div(num, bigIntBase)
		result = string(characterSet[remainder.Uint64()]) + result
	}

	return result
}
