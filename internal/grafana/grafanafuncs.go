package grafana

import (
	"time"
	"log/slog"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promauto"

	sim "sfop/internal/sim"

)



var partsProduced = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "parts_produced",
		Help: "Shows how many parts produced",
	},
	[]string{"machine", "namespace"},
)

var machineMaintenance = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "maintenanceCheck",
		Help: "If Machine is in maintenance",
	},
	[]string{"machine", "namespace"},
)

var totalOutput = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "totalPartsProduced",
		Help: "All Parts Produced in total",
	},
	[]string{"name", "namespace"},
)

var tempMachine = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "temp_per_machine",
		Help: "Shows temperature of machine",
	},
	[]string{"machine", "namespace"},
)

func updateMetrics() {
	for _, m := range sim.CncMachine {
		totalOutput.WithLabelValues("TotalPartsProduced", "Total Parts").Set(float64(sim.TotalParts))
		switch m.Status{
		case "running":
			machineMaintenance.DeleteLabelValues(m.MachineID, "maintenance")
			partsProduced.WithLabelValues(m.MachineID, m.Status).Set(float64(m.PartsProduced))
		    tempMachine.WithLabelValues(m.MachineID, m.Status).Set(float64(m.Temperature))
			
			
		case "maintenance":
            machineMaintenance.WithLabelValues(m.MachineID, m.Status).Set(1)		
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
