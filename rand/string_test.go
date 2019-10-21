package rand

import (
	"fmt"
	"testing"
)

func TestRandStringBytesMaskImprSrcUnsafe(t *testing.T) {
	str := GenRandString(10)
	fmt.Println(str)
}
