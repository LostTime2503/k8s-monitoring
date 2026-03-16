package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// Definerer en Counter-metrikk
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "myapp_http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint"},
	)

	// Definerer en Gauge-metrikk
	activeUsers = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "myapp_active_users",
			Help: "Current number of active users (simulated)",
		},
	)

	// Definerer et Histogram for responstid
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "myapp_request_duration_seconds",
			Help:    "Time spent processing requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"endpoint"},
	)
)

func init() {
	// Registrerer metrikkene hos Prometheus
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(activeUsers)
	prometheus.MustRegister(requestDuration)
}

// Simulerer aktive brukere
func simulateTraffic() {
	for {
		// Endrer antall aktive brukere tilfeldig
		activeUsers.Set(float64(rand.Intn(100)))
		time.Sleep(5 * time.Second)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// Simulerer litt arbeid
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)

	httpRequestsTotal.WithLabelValues(r.Method, "/").Inc()
	duration := time.Since(start).Seconds()
	requestDuration.WithLabelValues("/").Observe(duration)

	fmt.Fprintf(w, "Hello! Active users: %v", rand.Intn(100))
}

func main() {
	// Start simulering i bakgrunnen
	go simulateTraffic()

	http.HandleFunc("/", homeHandler)

	// Eksponerer /metrics endepunktet for Prometheus
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}