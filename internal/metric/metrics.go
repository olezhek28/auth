package metric

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

const appPrefix = "auth_server_"

type Metrics struct {
	requestTotal *prometheus.CounterVec
}

var metrics *Metrics

func Init(_ context.Context) error {
	metrics = &Metrics{
		requestTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: appPrefix + "requests_total",
				Help: "Количество запросов к серверу",
			},
			[]string{"method", "code"},
		),
	}

	return prometheus.Register(metrics.requestTotal)
}

func IncRequestTotal(method string, code string) {
	metrics.requestTotal.WithLabelValues(method, code).Inc()
}
