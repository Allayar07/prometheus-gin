package handler

import (
	"gin_prometheus/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type Handler struct {
	Services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{Services: service}
}

type Metric struct {
	MetricCollector prometheus.Collector
	ID              string
	Name            string
	Description     string
	Type            string
	Args            []string
}

var reqHistogram = &Metric{
	ID:          "reqHist",
	Name:        "requests_histogram_info",
	Description: "the latency of HTTP requests processed",
	Type:        "histogram_vec",
	Args:        []string{"URL", "status"},
}

type Prometheus struct {
	reqHistogram  *prometheus.HistogramVec
	router        *gin.Engine
	listenAddress string
	Metric        *Metric
	MetricsPath   string
}

func (h *Handler) InitRoutes() *gin.Engine {
	app := gin.Default()
	p := newPrometheus("http")
	p.use(app)
	app.POST("/create", h.Create)
	app.GET("/say", h.SaySomething)
	app.GET("/just", h.JustDoIt)
	return app
}

func newPrometheus(subsystem string) *Prometheus {
	p := &Prometheus{
		Metric:        reqHistogram,
		MetricsPath:   "/metrics",
		listenAddress: ":9911",
	}
	p.registerMetrics(subsystem)
	p.router = gin.Default()
	return p
}

func (p *Prometheus) registerMetrics(subsystem string) {
	metric := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "services",
			Name:      "request_duration_second",
			Help:      "Duration of the request",
			//Buckets:   prometheus.ExponentialBuckets(1, 1.3, 15), //50*1.3,15times
			Buckets: []float64{0.1, 0.15, 0.2, 0.25, 0.3},
		},
		reqHistogram.Args,
	)
	if err := prometheus.Register(metric); err != nil {
		logrus.Infof("%s could not be registered: %s", reqHistogram, err)
	} else {
		logrus.Infof("%s registered.", reqHistogram)
	}
	p.reqHistogram = metric
	reqHistogram.MetricCollector = metric
}

func (p *Prometheus) use(e *gin.Engine) {
	e.Use(p.handlerFunc())
	p.setMetricsPath(e)
}

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
		p.reqHistogram.WithLabelValues(c.Request.Host+c.Request.URL.String(), status).Observe(finished)
	}
}

func (p *Prometheus) setMetricsPath(e *gin.Engine) {
	p.router.GET(p.MetricsPath, prometheusHandler())
	go p.router.Run(p.listenAddress)
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
