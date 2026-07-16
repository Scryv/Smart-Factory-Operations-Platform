//make it so the whole operation sim is a function that can just be called in main
// get api work probs Chi or just basic net/http idk yet may try chi
// Prometheus grafana and Postgresql also slog for logging now postgres for maintenance machines etc and prom for datapoints like speed parts temps :)
// SLOG MUY GOAT

package main

import (
	"sfop/internal/sim"
)

func main() {
	sim.MachineSimulator()
}
