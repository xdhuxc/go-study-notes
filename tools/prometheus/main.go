package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var webRequestTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "web_request_total",
		Help: "Number of hello requests in total",
	},
	[]string{"method", "endpoint"},
)

var webRequestDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "web_request_duration_seconds",
		Help:    "web request duration distribution",
		Buckets: []float64{0.1, 0.3, 0.5, 0.7, 0.9, 2},
	},
	[]string{"method", "endpoint"},
)

func init() {
	// 注册监控指标
	prometheus.MustRegister(webRequestTotal)
	prometheus.MustRegister(webRequestDuration)
}

// 包装 handler function，不侵入业务逻辑
func Monitor(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h(w, r)
		duration := time.Since(start)
		// Counter 类型 metrics 的记录方法
		webRequestTotal.With(prometheus.Labels{"method": r.Method, "endpoint": r.URL.Path}).Inc()
		// Histogram 类型 metrics 的记录方式
		webRequestDuration.With(prometheus.Labels{"method": r.Method, "endpoint": r.URL.Path}).Observe(duration.Seconds())
	}
}

func Query(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	_, _ = io.WriteString(w, "some results")
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/query", Monitor(Query))
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
