package sim

import (
	"log/slog"
	"time"
)

type robotArms struct {
	MachineID        string
	Status           string //RUN MANTINANS REBot
	PartsMoved       int
	Temperature      int
	PowerConsumption float32
	Vibration        float32

	MaintenanceTime   time.Time
	MaintenanceReason string
}

func ProductGenARM(m *robotArms) {
	//RandomFloatingNumber := rand.Float64()
	//fmt.Println(RandomFloatingNumber)

	switch m.Status {
	case "running":

	case "maintenance":

	default:
		slog.Warn("not a valid status found for ",
			"machine id: ", m.MachineID)
	}

}
