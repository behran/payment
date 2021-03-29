package demon

import (
	"context"
	"time"

	"payment/internal/config"
	"payment/internal/demon/events"
	logger "payment/pkg/log"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

//Demon ...
type Demon struct {
	stop   chan struct{}
	ticker *time.Ticker
	events map[int]IEvent
}

//New ...
func New(config config.Config) *Demon {
	return &Demon{
		ticker: time.NewTicker(time.Duration(config.App.TimeRollback) * time.Second),
		stop:   make(chan struct{}),
		events: map[int]IEvent{
			RollBackEvent: events.NewRollbackTx(),
		},
	}
}

//Start Demon ...
func Start(lc fx.Lifecycle, demon *Demon) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			// start event ...
			go demon.Tick()
			return nil
		},
		OnStop: func(context.Context) error {
			// start event ...
			demon.Stop()
			return nil
		},
	})
}

//Tick ...
func (d Demon) Tick() {
	for {
		select {
		case <-d.ticker.C:
			if err := d.events[RollBackEvent].Run(); err != nil {
				logger.Logger.Error("error rollback event", zap.Error(err))
			}
		case <-d.stop:
			logger.Logger.Info("close rollback event")
			return
		}
	}
}

//Close ...
func (d Demon) Stop() { d.stop <- struct{}{} }
