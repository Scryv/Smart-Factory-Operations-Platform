package grafana

import (
	"time"
	"log/slog"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promauto"

	sim "sfop/internal/sim"

)



var cpuGauge = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "machine_cpu_percent",
		Help: "CPU usage of a machine",
	},
	[]string{"machine", "namespace"},
)

var extra = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "dsdsds",
		Help: "CPU usage of a",
	},
	[]string{"machine", "namespace"},
)

var memGauge = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "machine_memory_mb",
		Help: "Memory usage of a machine",
	},
	[]string{"machine", "namespace"},
)

func updateMetrics() {
	for _, m := range sim.CncMachine {
		switch m.Status{
		case "running":
			extra.DeleteLabelValues(m.MachineID, "maintenance")
			cpuGauge.WithLabelValues(m.MachineID, m.Status).Set(float64(m.PartsProduced))
		    memGauge.WithLabelValues(m.MachineID, m.Status).Set(float64(m.Temperature))
			
		case "maintenance":
            extra.WithLabelValues(m.MachineID, m.Status).Set(1)		
		default:
			slog.Error("No clue why this is in maintenance?")
		
		}
	}
}
func bf() {
	prometheus.Unregister(collectors.NewGoCollector())
}
func GrafanaMetUpdate() {
	bf()

	// Simulate metrics changing every 5 seconds
	go func() {
		for {
			updateMetrics()
			time.Sleep(1 * time.Second)
		}
	}()
}
