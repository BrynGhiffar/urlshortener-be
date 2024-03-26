package main

import (
	"math/rand"
	"os"
	"strconv"
	"time"
)

const poss = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateRandomString(length int) string {
	res := make([]byte, length)
	for i := range res {
		res[i] = poss[rand.Intn(len(poss))]
	}
	return string(res)
}

func getRedirectExpirationSecsEnv() int64 {
	redirectExpirationEnv, found := os.LookupEnv("REDIRECT_EXPIRATION_SECS")
	if !found {
		panic("REDIRECT_EXPIRATION_SECS environment variable not defined")
	}
	if val, err := strconv.ParseInt(redirectExpirationEnv, 10, 64); err != nil {
		panic("REDIRECT_EXPIRATION_SECS environment variable parsing error")
	} else {
		return val
	}
}

func getSwaggerHostnameEnv() string {
	hostnameEnv, found := os.LookupEnv("SWAGGER_HOSTNAME")
	if !found {
		panic("SWAGGER_HOSTNAME environment variable not defined")
	}
	return hostnameEnv
}

func expiresAfter(secs int64) int64 {
	return time.Now().UTC().Unix() + secs
}
