package influxstore

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	_ "github.com/influxdata/influxdb-client-go/v2/api/write"
)

type InfluxStore struct {
	influxClient influxdb2.Client
}

func New(client influxdb2.Client) *InfluxStore {
	return &InfluxStore{
		influxClient: client,
	}
}
