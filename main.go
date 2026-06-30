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
	Temperature      float32
	PowerConsumption float32
	Vibration        float32
	SpindleSpeed     int

	MaintenanceTime time.Time
}

type robotArms struct {
	MachineID        string
	Status           string //RUN MANTINANS REBot
	PartsMoved       int
	Temperature      float32
	PowerConsumption float32
	Vibration        float32

	MaintenanceTime time.Time
}

var maintenanceLogs []MaintenanceTab

func productGenCNC(m *CNC) {
	//RandomFloatingNumber := rand.Float64()
	//fmt.Println(RandomFloatingNumber)
	currentTime := time.Now()
	switch m.Status {
	case "running":
		MaintenanceTest := MaintenanceTab{
			MachineID:           m.MachineID,
			MaintenanceStart:    currentTime,
			MaintenanceEnd:      currentTime,
			MaintenanceDuration: currentTime,
			Reason:              "bc fuh nig",
		}
		maintenanceLogs = append(maintenanceLogs, MaintenanceTest)
		fmt.Println(maintenanceLogs)
	case "maintenance":

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
		MachineID:        "CNC-1",
		Status:           "running",
		PartsProduced:    0,
		Temperature:      0.0,
		PowerConsumption: 0.0,
		Vibration:        0.0,
		SpindleSpeed:     0,
	}
	robotArm := robotArms{
		MachineID:        "RobotARM-1",
		Status:           "running",
		PartsMoved:       0,
		Temperature:      0.0,
		PowerConsumption: 0.0,
		Vibration:        0.0,
		MaintenanceTime:  currentTime,
	}
	cnc1 := CNC{
		MachineID:        "CNC-2",
		Status:           "running",
		PartsProduced:    0,
		Temperature:      0.0,
		PowerConsumption: 0.0,
		Vibration:        0.0,
		MaintenanceTime:  currentTime,
		SpindleSpeed:     0,
	}
	robotArm1 := robotArms{
		MachineID:        "RobotARM-2",
		Status:           "running",
		PartsMoved:       0,
		Temperature:      0.0,
		PowerConsumption: 0.0,
		Vibration:        0.0,
		MaintenanceTime:  currentTime,
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
