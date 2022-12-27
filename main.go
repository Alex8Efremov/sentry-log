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
	message  = os.Getenv("MESSAGE")
)

func main() {
	var test int = 0
	sentrySyncTransport := sentry.NewHTTPSyncTransport()
	sentrySyncTransport.Timeout = time.Second * 1

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              dsn,
		Release:          "my-project-name@1.0.0",
		Environment:      "production",
		Debug:            true,
		TracesSampleRate: 1.0,
		Transport:        sentrySyncTransport,
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
		allert := fmt.Sprintf("%s, Allert number: %d", message, test)
		sentry.CaptureMessage(allert)
		time.Sleep(1 * time.Millisecond)
		if test == i {
			break
		}
		test++
		fmt.Println(test)

	}
	fmt.Println("End!")
}
