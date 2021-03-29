package facade

import "payment/internal/factory"

//Service ...
func Service() *factory.ServiceFactory { return factory.Services }
