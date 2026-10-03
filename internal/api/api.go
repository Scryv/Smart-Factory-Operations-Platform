package api

import (
	"net/http"
	"log"
    "github.com/prometheus/client_golang/prometheus/promhttp"
	gf "sfop/internal/grafana"
)

func Sapi() {
	gf.GrafanaMetUpdate()
	http.Handle("/metrics", promhttp.Handler())
	log.Println("Listening on :5701")
	log.Fatal(http.ListenAndServe(":5701", nil))
}
