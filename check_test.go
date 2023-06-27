package cci

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	info, ok, err := Check("110101199901011232") // 有效
	// info, ok, err := Check("110101199901011234") // 无效
	if err != nil {
		panic(err)
	}
	fmt.Println("是否有效", ok)
	fmt.Println(info)
}
