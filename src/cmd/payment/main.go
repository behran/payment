package main

import (
	"context"
	"log"

	application "payment/internal/app"
)

func main() {
	ctx := context.Background()

	containers := application.Containers()

	if err := containers.Start(ctx); err != nil {
		log.Fatalf("Application didn't start `err`:%s", err)
	}
	<-containers.Done()

	if err := containers.Stop(ctx); err != nil {
		log.Fatalf("Application didn't stop `err`:%s", err)
	}
}
