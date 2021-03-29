package main

import (
	"time"

	application "payment/internal/app"
	logger "payment/pkg/log"

	"go.uber.org/zap"
)

func main() {
	// waiting postgres container ...
	time.Sleep(20*time.Second)
	if err := application.Run(); err != nil {
		logger.Logger.Fatal("Application didn't start",
			zap.Error(err),
		)
	}
}
