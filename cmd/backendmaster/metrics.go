package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	writeEvent = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "",
		Name:      "writeFile",
		Help:      "Number of writes made to file",
	},
		[]string{"filename"},
	)

// serviceRestart = prometheus.NewCounterVec(
// 	prometheus.CounterOpts{
// 		Namespace: "puppeteer",
// 		Name:      "service_restart",
// 		Help:      "Service Restart Process Count",
// 	},
// 	[]string{"service", "status"},
// )
)
