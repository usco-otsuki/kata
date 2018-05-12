package easy_diagonal

import "math/big"

func Diagonal(n, p int) int {
	return int((&big.Int{}).Binomial(int64(n+1), int64(p+1)).Int64())
}
