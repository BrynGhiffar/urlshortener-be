package main

import (
	"time"

	"github.com/robfig/cron"
)

func main() {
	r := setupRouter()
	c := cron.New()
	c.AddFunc("@every 5s", func() {
		toDelete := []string{}
		now := time.Now().UTC().Unix()
		for key, redirect := range db {
			if now > redirect.expiration {
				toDelete = append(toDelete, key)
			}
		}
		for _, key := range toDelete {
			delete(db, key)
		}
	})
	c.Start()
	r.Run(":8080")
	c.Stop()
}
