package main

import (
	autoconfix "booqall.com/libraries/autoconfix/pkg"
	runtmx "booqall.com/libraries/runtmx/pkg"
	nats "github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

func initializers() []autoconfix.Initializer {
	initializers := make([]autoconfix.Initializer, 0)

	// To Do

	return initializers
}

func services() []nats.Subscription {
	services := make([]nats.Subscription, 0)

	// To Do

	return services
}

func init() {
	autoconfix.Bootstrap(initializers()...)
}

func main() {
	services := services()
	logrus.Info("Service Started")
	runtmx.AwaitInterrupt(func() {
		logrus.Warn("Service is Shutting Down")

		for _, service := range services {
			service.Unsubscribe()
		}

		logrus.Warn("Service Shut Down")
	})
}
