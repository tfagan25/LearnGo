package main

import (
	"time"
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

func recordGauges() {
	go func() {
		for {
			testGauge.Set(0)
			time.Sleep(2 * time.Second)
			testGauge.Set(1)
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
			Name: "myapp_processed_ops_total",
			Help: "The total number of processed events",
	})
)

var (
	testGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "my_test_gauge",
		Help: "Testing a gauge metric",
	})
)

func main () {
	recordMetrics()
	recordGauges()
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":3000", nil)
}