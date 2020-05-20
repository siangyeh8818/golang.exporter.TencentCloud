package server

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Exporter struct {
	account_balance prometheus.Gauge
	//gaugeVec prometheus.GaugeVec
}

func Run_Exporter_Server() {
	log.Println(`
  This is a prometheus exporter for Tencentcloud
  Access: http://127.0.0.1:8081
  `)

	metricsPath := "/metrics"
	listenAddress := ":8081"
	metricsPrefix := "tencentcloud"
	exporters := NewExporter(metricsPrefix)
	/*
	   	registry := prometheus.NewRegistry()
	       registry.MustRegister(metrics)
	*/
	prometheus.MustRegister(exporters)

	// Launch http service

	http.Handle(metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
		 <head><title>Dummy Exporter</title></head>
		 <body>
		 <h1>Dummy Exporter</h1>
		 <p><a href='` + metricsPath + `'>Metrics</a></p>
		 </body>
		 </html>`))
	})
	log.Println(http.ListenAndServe(listenAddress, nil))
}

func NewExporter(metricsPrefix string) *Exporter {
	gauge_metrics1 := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:      "account_balance",
		Help:      "This is a gauge metric example"})

	/*
		gaugeVec := *prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: metricsPrefix,
			Name:      "gauge_vec_metric",
			Help:      "This is a siang gauga vece metric"},
			[]string{"myLabel"})
	*/
	return &Exporter{
		account_balance: gauge_metrics1,
		//gaugeVec: gaugeVec,
	}
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	//e.gauge.Set(float64(100))

	e.account_balance.Set(callAPI("AccountBalance"))
	//e.gaugeVec.WithLabelValues("hello").Set(float64(0))
	e.account_balance.Collect(ch)
	//e.gaugeVec.Collect(ch)
}

// 讓exporter的prometheus屬性呼叫Describe方法

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	e.account_balance.Describe(ch)
}
