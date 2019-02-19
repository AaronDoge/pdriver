package cron

import (
	"fmt"
	"github.com/robfig/cron"
	"math/rand"
	"os"
	"strings"
	"time"
)

func SetCron() error {

	c := cron.New()
	err := c.AddFunc("*/3 * * * * ?", func() {
		fd, err := os.OpenFile("../data/cron.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
		defer fd.Close()
		if err != nil {
			return
		}
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		id := RandStringRunes(16)

		content := strings.Join([]string{"[INFO][CRON] ", currentTime, "  ", id, "   <under tracing> \n"}, "")
		fd.Write([]byte(content))

	})
	c.Start()

	if err != nil {
		return err
	}

	fmt.Println("启动crontab成功。。。")

	return nil

}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%&*")


func RandStringRunes(n int) string {

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

