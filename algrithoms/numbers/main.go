package main

import (
	"github.com/civet148/gotools/log"
	"math"
)

func main() {
	DecNumberTruncateLast(123456789, 5)
}

//十进制数取后面任意n位值
func DecNumberTruncateLast(number, n int) {
	var d int
	log.Debugf("number=%d n=%d", number, n)
	d = number % (int(math.Pow10(n)))
	log.Debugf("number%%(Pow(10,n))=%d", d)
	return
}
