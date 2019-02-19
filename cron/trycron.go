package cron

import (
	"github.com/robfig/cron"
	"log"
	"os"
	"strings"
	"time"
)

func TryCron() {
	i := 0
	c := cron.New()
	spec := "*/5 * * * * ?"
	c.AddFunc(spec, func() {
		fd, err := os.OpenFile("../data/cron.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		defer fd.Close()
		if err != nil {
			return
		}
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		id := RandStringRunes(16)

		content := strings.Join([]string{"[INFO][CRON] ", currentTime, "  ", id, "   <under tracing> \n"}, "")
		i++
		log.Println("cron running:", i, ">>>", content)
		fd.Write([]byte(content))
	})
	c.Start()

	select {}
}
