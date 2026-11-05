package sim




var (
	ArmMachine = []*robotArms{
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
)