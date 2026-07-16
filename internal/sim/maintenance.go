package sim

import (
	"time"
)

const (
	ovrHeat string = "Overheating"
	bogDown string = "Bog Down"
	highVib string = "Vibrations Too High"
)

type MaintenanceTab struct {
	MachineID           string
	MaintenanceStart    time.Time
	MaintenanceEnd      time.Time
	MaintenanceDuration time.Duration
	Reason              string
}

var maintenanceLogs []MaintenanceTab
