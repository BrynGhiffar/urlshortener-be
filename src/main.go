package main

func main() {
	r := setupRouter()
	c := setupCron()
	c.Start()
	r.Run(":8080")
}
