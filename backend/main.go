package main

import (
	"backend/services"
	"backend/system/actions"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/cpu"
	"log"
	"runtime"
	"time"
)

func main() {
	services.InitDBConnection()

	r := gin.Default()
	routes(r)

	r.Run(":80")
}

func routes(r *gin.Engine) {
	r.GET("/ping", actions.Ping)

	prometheusInit(r)
}

func prometheusInit(r *gin.Engine) {
	myMetric := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "backend_api",
		Help: "API Metrics",
	})

	ramUsage := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "backend_api_ram_usage_bytes",
		Help: "RAM usage of the API service",
	})

	cpuUsage := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "backend_api_cpu_usage_percent",
		Help: "Current CPU usage percent of API Service",
	})

	prometheus.MustRegister(myMetric, ramUsage, cpuUsage)

	go func() {
		for {
			percent, err := cpu.Percent(time.Second, false)
			if err != nil {
				log.Println("Ошибка при получении процента использования ЦПУ:", err)
				continue
			}
			cpuUsage.Set(percent[0])
			time.Sleep(10 * time.Second)
		}
	}()

	go func() {
		for {
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			ramUsage.Set(float64(m.Alloc) / 1024 / 1024)
			time.Sleep(10 * time.Second)
		}
	}()

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
