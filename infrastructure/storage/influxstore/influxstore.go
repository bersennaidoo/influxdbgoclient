package influxstore

import (
	"context"
	"log"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

type InfluxStore struct {
	client influxdb2.Client
}

func New(client influxdb2.Client) *InfluxStore {
	return &InfluxStore{
		client: client,
	}
}

func (inf *InfluxStore) WriteDataPoint(org, bucket string) {
	writeAPI := inf.client.WriteAPIBlocking(org, bucket)
	for value := 0; value < 5; value++ {
		tags := map[string]string{
			"tagname1": "tagvalue1",
		}
		fields := map[string]interface{}{
			"field1": value,
		}
		point := write.NewPoint("measurement1", tags, fields, time.Now())
		time.Sleep(1 * time.Second) // separate points by 1 second

		if err := writeAPI.WritePoint(context.Background(), point); err != nil {
			log.Fatal(err)
		}
		log.Printf("Data Point stored in bucket")
	}
	log.Printf("Done")
}
