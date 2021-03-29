package factory

import (
	"sync"

	"payment/internal/database"
	"payment/internal/domains/payment"
	"payment/internal/domains/payment/repositories"

	"golang.org/x/sync/errgroup"
)

//ServiceFactory ...
type ServiceFactory struct {
	payment   *payment.Service
	connector *database.ConnectManager
}

var (
	Services *ServiceFactory
	once     sync.Once
)

//NewFactory ...
func NewFactory(connector *database.ConnectManager) *ServiceFactory {
	once.Do(func() {
		Services = &ServiceFactory{
			connector: connector,
		}
	})
	return Services
}

//InitServices ...
func InitServices(factory *ServiceFactory) error {
	var eg errgroup.Group

	for _, service := range []func() error{
		factory.createPayment,
	} {
		eg.Go(service)
	}
	return eg.Wait()
}

func (factory *ServiceFactory) createPayment() error {
	repository, err := repositories.NewAccount(factory.connector)
	if err != nil {
		return err
	}
	factory.payment = payment.New(
		repository,
	)
	return nil
}

//Payment ...
func (factory ServiceFactory) Payment() *payment.Service { return factory.payment }
