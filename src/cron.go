package main

import (
	"time"

	"github.com/robfig/cron"
)

func removeExpiredRedirects() {
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
}

func setupCron() *cron.Cron {
	c := cron.New()
	c.AddFunc("@every 5s", removeExpiredRedirects)
	return c
}
