package metrics

import (
	"gin_prometheus/internal/helpers"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"strconv"
	"time"
)

type Prometheus struct {
	reqDuration   *prometheus.HistogramVec
	appVersion    *prometheus.GaugeVec
	router        *gin.Engine
	listenAddress string
	MetricsPath   string
}

func NewPrometheus() *Prometheus {
	p := &Prometheus{
		MetricsPath:   "/metrics",
		listenAddress: ":9911",
	}
	p.registerMetrics()
	p.router = gin.Default()
	return p
}

func (p *Prometheus) registerMetrics() {
	HistogramMetric := helpers.NewHistogramVecReqLatency()
	GaugeMetric := helpers.NewGaugeVersion()
	p.reqDuration = HistogramMetric
	p.appVersion = GaugeMetric
}

// Middleware to set histogram metric for all requests
func (p *Prometheus) handlerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		if c.Request.URL.String() == p.MetricsPath {
			c.Next()
			return
		}
		c.Next()

		status := strconv.Itoa(c.Writer.Status())
		finished := time.Since(start).Seconds()
		p.reqDuration.WithLabelValues(c.Request.Host+c.Request.URL.String(), status).Observe(finished)
	}
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (p *Prometheus) setMetricsPath(e *gin.Engine) {
	p.router.GET(p.MetricsPath, prometheusHandler())
	go p.router.Run(p.listenAddress)
}

func (p *Prometheus) Use(e *gin.Engine) {
	e.Use(p.handlerFunc())
	p.setMetricsPath(e)
}
