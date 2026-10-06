package goutils

import (
	"log"
	"os"
	"time"
)

func GetEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}

func MustGetEnv(key string) string {
	val := GetEnv(key, "")

	if val == "" {
		log.Println("Value for this key: ", key, " is not in env")
	}
	return val
}

func NowIST(t time.Time) time.Time {
	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		loc = time.FixedZone("IST", 5*3600+1800)
	}
	return t.In(loc)
}
