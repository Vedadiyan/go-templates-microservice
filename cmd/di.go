package main

import (
	connex "booqall.com/libraries/connex/pkg"
	iocx "booqall.com/libraries/iocx/pkg"
	natx "booqall.com/pipelines/nats/pkg/encoded"
	"github.com/nats-io/nats.go"
)

func AddNats(addr string) {
	conn := natx.NatsConnection{
		ConnectionProperties: connex.ConnectionProperties{
			Endpoints: []string{
				addr,
			},
		},
	}
	conn.Create()
	iocx.AddSinleton(func() (*nats.EncodedConn, error) {
		natsConn, _, err := conn.Get()
		return natsConn, err
	})
}
