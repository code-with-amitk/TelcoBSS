package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	Requests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "telcobss_requests_total",
			Help: "Total number of API requests processed.",
		},
		[]string{"service", "endpoint", "status"},
	)
)

func InitPrometheus() {
	prometheus.MustRegister(Requests)
}
