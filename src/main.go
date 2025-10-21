package main

import "os"

func main() {
	r := setupRouter()
	c := setupCron()
	c.Start()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
