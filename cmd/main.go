package main

import (
	"os"

	"github.com/bersennaidoo/influxdbgoclient/physical/influxd"
)

func main() {
	token := os.Getenv("INFLUXDB_TOKEN")
	url := "http://localhost:8086"

	client := influxd.CreateInfluxDBClient(url, token)

	inflst := influxstore.New(client)
}
