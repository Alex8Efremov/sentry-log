package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/getsentry/sentry-go"
)

var (
	dsn      = os.Getenv("DSN")
	attempts = os.Getenv("ATTEMPTS")
)

func main() {
	var test int = 0
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              dsn,
		Release:          "my-project-name@1.0.0",
		Environment:      "production",
		Debug:            true,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)
	defer sentry.CaptureMessage("It's a defer!")

	i, err := strconv.Atoi(attempts)
	if err != nil {
		// ... handle error
		panic(err)
	}

	for {
		sentry.CaptureMessage("It works!")
		time.Sleep(1 * time.Second)
		if test == i {
			break
		}
		test++
		fmt.Println(test)

	}
	fmt.Println("End!")
}
