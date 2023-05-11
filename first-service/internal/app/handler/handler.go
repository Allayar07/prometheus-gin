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

var reqCnt = &Metric{
	ID:          "reqCnt",
	Name:        "requests_total",
	Description: "the number of HTTP requests processed",
	Type:        "counter_vec",
	Args:        []string{"status"}}

type Prometheus struct {
	reqCnt        *prometheus.SummaryVec
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
	//app.GET("/metrics", gin.WrapH(promhttp.Handler()))
	return app
}

func newPrometheus(subsystem string) *Prometheus {
	p := &Prometheus{
		Metric:        reqCnt,
		MetricsPath:   "/metrics",
		listenAddress: ":9911"}
	p.registerMetrics(subsystem)
	p.router = gin.Default()
	return p
}

func (p *Prometheus) registerMetrics(subsystem string) {
	metric := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace:  "services",
			Subsystem:  "http",
			Name:       "request",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		reqCnt.Args,
	)
	if err := prometheus.Register(metric); err != nil {
		logrus.Infof("%s could not be registered: ", reqCnt, err)
	} else {
		logrus.Infof("%s registered.", reqCnt)
	}
	p.reqCnt = metric

	reqCnt.MetricCollector = metric
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
		p.reqCnt.WithLabelValues(status).Observe(finished)
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
