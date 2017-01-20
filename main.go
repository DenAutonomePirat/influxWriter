package main

import (
	"github.com/influxdata/influxdb/client/v2"
	"math/rand"
	"time"
)

func main() {
	config := client.UDPConfig{
		Addr: "10.0.0.1:8089",
	}
	// Make client
	c, err := client.NewUDPClient(config)
	if err != nil {
		panic(err.Error())
	}

	for {
		r := rand.New(rand.NewSource(99))
		val := r.Float64()
		WriteUDP(val)
		time.Sleep(1 * time.Second)
	}

}

func WriteUDP(r float64) {
	// Create a new point batch
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Precision: "u",
	})

	// Create a point and add to batch
	tags := map[string]string{"cpu": "cpu-total"}
	fields := map[string]interface{}{
		"idle":   r,
		"system": 53.3,
		"user":   46.6,
	}
	pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
	if err != nil {
		panic(err.Error())
	}
	bp.AddPoint(pt)

	// Write the batch
	c.Write(bp)
}
