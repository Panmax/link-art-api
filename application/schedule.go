package application

import (
	"github.com/robfig/cron/v3"
	"log"
)

// https://godoc.org/github.com/robfig/cron
func SetupSchedule() {
	go func() {
		cronTab := cron.New()

		cronTab.AddFunc("*/1 * * * *", func() {
			log.Println("1111111")
		})

		cronTab.AddFunc("*/2 * * * *", func() {
			log.Println("222222")
		})

		cronTab.Start()
	}()
}
