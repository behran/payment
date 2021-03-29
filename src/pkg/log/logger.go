package logger

import (
	"log"
	"sync"

	"go.uber.org/zap"
)

//Logger ...
var (
	Logger *zap.Logger
	once   sync.Once
)

//NewLogger ...
func New() *zap.Logger {
	once.Do(func() {
		var err error
		Logger, err = zap.NewProduction()
		if err != nil {
			log.Fatal(err)
		}
	})
	return Logger
}
