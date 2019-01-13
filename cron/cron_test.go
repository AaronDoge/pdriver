package cron

import (
	"fmt"
	"testing"
)

func TestSetCron(t *testing.T) {
	if err := SetCron(); err != nil {
		fmt.Println("启动cron失败, ", err.Error())
	}
}


func TestRandStringRunes(t *testing.T) {
	s := RandStringRunes(10)
	fmt.Println("生成的随机字符串：", s)
}