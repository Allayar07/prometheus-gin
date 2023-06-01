package helpers

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

func NewHistogramVecReqLatency() *prometheus.HistogramVec {
	HistogramMetric := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "practice_service",
			Name:      "request_duration_seconds",
			Help:      "Duration of the request.",
			// 4 times larger for apdex score
			//Buckets: prometheus.ExponentialBuckets(0.1, 1.5, 5),
			//Buckets: prometheus.LinearBuckets(0.01, 0.05, 5),
			Buckets: []float64{0.1, 0.15, 0.2, 0.25, 0.3},
		},
		[]string{"URL", "status"},
	)
	if err := prometheus.Register(HistogramMetric); err != nil {
		logrus.Infof("%s could not be registered: %s", "HistogramVec type", err)
	} else {
		logrus.Infof("%s registered.", "HistogramVec type")
	}
	return HistogramMetric
}

func NewGaugeVersion() *prometheus.GaugeVec {
	GaugeMetric := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "myapp",
			Name:      "info",
			Help:      "Information about the My App version",
		},
		[]string{"version"})
	GaugeMetric.With(prometheus.Labels{"version": "v.0.1"}).Set(1)
	if err := prometheus.Register(GaugeMetric); err != nil {
		logrus.Infof("%s could not be registered: %s", "GaugeVec type", err)
	} else {
		logrus.Infof("%s registered.", "GaugeVec type")
	}
	return GaugeMetric
}
