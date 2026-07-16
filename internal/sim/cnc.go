package sim

import (
	"log/slog"
	"math"
	"math/rand/v2"
	"time"
)

type CNC struct {
	MachineID          string
	Status             string //RUN MANTINANS REBot
	PartsProduced      int    //2 to 5 but penalty when things dont doe what suposed to do
	Temperature        int
	PowerConsumption   float32
	Vibration          float32
	SpindleSpeed       float64
	TargetSpindleSpeed float64

	MaintenanceTime   time.Time
	MaintenanceReason string
}

func ProductGenCNC(m *CNC) {
	//RandomFloatingNumber := rand.Float64()
	//fmt.Println(RandomFloatingNumber)

	switch m.Status { //add bog down if under 15% of expected RPM
	case "running":
		currentTime := time.Now()
		bogger := rand.Float64()
		if bogger < 0.01 {
			m.SpindleSpeed = 0.15 * m.TargetSpindleSpeed
			slog.Error(bogDown)
			m.Status = "maintenance"
			m.MaintenanceReason = bogDown
			m.MaintenanceTime = currentTime
		}

		tempAdd := int(math.Round(m.TargetSpindleSpeed / m.SpindleSpeed))

		m.SpindleSpeed += rand.Float64()*24 - 14
		m.Temperature += tempAdd + (rand.IntN(2) + 1)

		if m.Temperature > 85 {
			slog.Error(ovrHeat)
			m.Status = "maintenance"
			m.MaintenanceReason = ovrHeat
			m.MaintenanceTime = currentTime
		}

		if m.SpindleSpeed < (0.10 * m.TargetSpindleSpeed) {
			if m.SpindleSpeed < (0.15 * m.TargetSpindleSpeed) {
				slog.Error(bogDown)
				m.Status = "maintenance"
				m.MaintenanceReason = bogDown
				m.MaintenanceTime = currentTime
			}
			m.PartsProduced += (rand.IntN(2) + 1)
		} else {
			m.PartsProduced += (rand.IntN(4) + 2)
		}

	case "maintenance":
		switch m.MaintenanceReason {
		case ovrHeat:
			if m.Temperature > 25 {
				m.Temperature -= (rand.IntN(4) + 2)
			} else {
				currentTime := time.Now()
				MaintenanceTest := MaintenanceTab{
					MachineID:           m.MachineID,
					MaintenanceStart:    m.MaintenanceTime,
					MaintenanceEnd:      currentTime,
					MaintenanceDuration: currentTime.Sub(m.MaintenanceTime),
					Reason:              m.MaintenanceReason,
				}
				maintenanceLogs = append(maintenanceLogs, MaintenanceTest)
				slog.Info("Overheating finished")
				m.Status = "running"
				m.MaintenanceReason = ""
			}
		case bogDown:
			if m.SpindleSpeed < 3200 {
				m.SpindleSpeed += float64(rand.IntN(24) + 3)
			} else {
				currentTime := time.Now()
				MaintenanceTest := MaintenanceTab{
					MachineID:           m.MachineID,
					MaintenanceStart:    m.MaintenanceTime,
					MaintenanceEnd:      currentTime,
					MaintenanceDuration: currentTime.Sub(m.MaintenanceTime),
					Reason:              m.MaintenanceReason,
				}
				maintenanceLogs = append(maintenanceLogs, MaintenanceTest)
				slog.Info("bogdown finished")
				m.Status = "running"
				m.MaintenanceReason = ""
			}
		case highVib:
		default:
			slog.Error("No clue why this is in maintenance?")
		}
	default:
		slog.Warn("not a valid status found for ",
			"machine id: ", m.MachineID)
	}

}
