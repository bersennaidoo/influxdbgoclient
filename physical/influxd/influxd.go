package influxd

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func CreateInfluxDBClient(url, token string) influxdb2.Client {
	client := influxdb2.NewClient(url, token)
	return client
}
