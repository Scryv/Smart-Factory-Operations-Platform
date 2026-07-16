package sim

import (
	"fmt"
	"time"
)

func MachineSimulator() {
	currentTime := time.Now()

	cncMachine := []*CNC{
		{
			MachineID:          "CNC-1",
			Status:             "running",
			PartsProduced:      0,
			Temperature:        0,
			PowerConsumption:   0.0,
			Vibration:          0.08,
			SpindleSpeed:       3200.0,
			TargetSpindleSpeed: 3500.0,
			MaintenanceReason:  "none",
		},
		{
			MachineID:          "CNC-2",
			Status:             "running",
			PartsProduced:      0,
			Temperature:        0,
			PowerConsumption:   0.0,
			Vibration:          0.08,
			MaintenanceTime:    currentTime,
			SpindleSpeed:       3200.0,
			TargetSpindleSpeed: 3500.0,
			MaintenanceReason:  "none",
		},
	}

	armMachine := []*robotArms{
		{
			MachineID:         "RobotARM-1",
			Status:            "running",
			PartsMoved:        0,
			Temperature:       0,
			PowerConsumption:  0.0,
			Vibration:         0.0,
			MaintenanceTime:   currentTime,
			MaintenanceReason: "none",
		},
		{
			MachineID:         "RobotARM-2",
			Status:            "running",
			PartsMoved:        0,
			Temperature:       0,
			PowerConsumption:  0.0,
			Vibration:         0.0,
			MaintenanceTime:   currentTime,
			MaintenanceReason: "none",
		},
	}

	for {

		for _, cncM := range cncMachine {
			fmt.Println(cncM.MachineID)
			ProductGenCNC(cncM)
			fmt.Println(cncM.PartsProduced)

		}
		for _, armM := range armMachine {
			fmt.Println(armM.MachineID)
			ProductGenARM(armM)
		}
		fmt.Println("Hello Cuh")
		time.Sleep(1000 * time.Millisecond)
	}
}
