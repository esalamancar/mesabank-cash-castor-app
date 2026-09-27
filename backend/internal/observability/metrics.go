package observability

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total de requests HTTP procesadas, por método, ruta y status.",
	}, []string{"method", "path", "status"})

	httpRequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_request_duration_seconds",
		Help: "Duración de las requests HTTP, por método y ruta.",
	}, []string{"method", "path"})

	// WebsocketConnections cuenta las conexiones WebSocket activas.
	WebsocketConnections = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "websocket_connections_active",
		Help: "Conexiones WebSocket activas en el hub.",
	})
)

// Metrics registra cada request HTTP en los contadores/histogramas de
// Prometheus.
func Metrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		path := c.FullPath()
		if path == "" {
			path = "unmatched"
		}
		status := strconv.Itoa(c.Writer.Status())

		httpRequestsTotal.WithLabelValues(c.Request.Method, path, status).Inc()
		httpRequestDuration.WithLabelValues(c.Request.Method, path).Observe(time.Since(start).Seconds())
	}
}

// Handler expone /metrics en formato Prometheus.
func Handler() gin.HandlerFunc {
	h := promhttp.Handler()
	return gin.WrapH(h)
}
