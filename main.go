package main

import (
	"fmt"
	"log"
	"time"
)

type MaintenanceTab struct {
	MachineID           string
	MaintenanceStart    time.Time
	MaintenanceEnd      time.Time
	MaintenanceDuration time.Time
	Reason              string
}

type CNC struct {
	MachineID        string
	Status           string //RUN MANTINANS REBot
	PartsProduced    int    //2 to 5 but penalty when things dont doe what suposed to do
	Temperature      int
	PowerConsumption float32
	Vibration        float32
	SpindleSpeed     int

	MaintenanceTime   time.Time
	MaintenanceReason string
}

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

var maintenanceLogs []MaintenanceTab

func productGenCNC(m *CNC) {
	//RandomFloatingNumber := rand.Float64()
	//fmt.Println(RandomFloatingNumber)
	currentTime := time.Now()
	switch m.Status {
	case "running":
		if m.Temperature > 55 {
			m.Status = "maintenance"
			m.MaintenanceReason = "Overheating"
		}

	case "maintenance":
		MaintenanceTest := MaintenanceTab{
			MachineID:           m.MachineID,
			MaintenanceStart:    currentTime,
			MaintenanceEnd:      currentTime,
			MaintenanceDuration: currentTime,
			Reason:              m.MaintenanceReason,
		}
		maintenanceLogs = append(maintenanceLogs, MaintenanceTest)
		fmt.Println(maintenanceLogs)

	default:
		log.Println("not a valid status found for ", m.MachineID)
	}

}

func productGenARM(m *robotArms) {
	//RandomFloatingNumber := rand.Float64()
	//fmt.Println(RandomFloatingNumber)
	switch m.Status {
	case "running":

	case "maintenance":

	default:
		log.Println("not a valid status found for ", m.MachineID)
	}

}
func main() {
	currentTime := time.Now()

	cnc := CNC{
		MachineID:         "CNC-1",
		Status:            "running",
		PartsProduced:     0,
		Temperature:       0,
		PowerConsumption:  0.0,
		Vibration:         0.0,
		SpindleSpeed:      0,
		MaintenanceReason: "none",
	}
	robotArm := robotArms{
		MachineID:         "RobotARM-1",
		Status:            "running",
		PartsMoved:        0,
		Temperature:       0,
		PowerConsumption:  0.0,
		Vibration:         0.0,
		MaintenanceTime:   currentTime,
		MaintenanceReason: "none",
	}
	cnc1 := CNC{
		MachineID:         "CNC-2",
		Status:            "running",
		PartsProduced:     0,
		Temperature:       0,
		PowerConsumption:  0.0,
		Vibration:         0.0,
		MaintenanceTime:   currentTime,
		SpindleSpeed:      0,
		MaintenanceReason: "none",
	}
	robotArm1 := robotArms{
		MachineID:         "RobotARM-2",
		Status:            "running",
		PartsMoved:        0,
		Temperature:       0,
		PowerConsumption:  0.0,
		Vibration:         0.0,
		MaintenanceTime:   currentTime,
		MaintenanceReason: "none",
	}

	cncMachine := []*CNC{
		&cnc,
		&cnc1,
	}

	armMachine := []*robotArms{
		&robotArm,
		&robotArm1,
	}

	for {

		for _, cncM := range cncMachine {
			fmt.Println(cncM.MachineID)
			productGenCNC(cncM)

		}
		for _, armM := range armMachine {
			fmt.Println(armM.MachineID)
			productGenARM(armM)
		}
		fmt.Println("Hello Cuh")
		time.Sleep(1000 * time.Millisecond)
	}
}
