package sim

import (
	"fmt"
	"time"

)
var currentTime = time.Now()
var TotalParts = 0


func MachineSimulator() {
	for {

		for _, cncM := range CncMachine {
			fmt.Println(cncM.MachineID)
			ProductGenCNC(cncM)
			fmt.Println(cncM.PartsProduced)

		}
		for _, armM := range ArmMachine {
			fmt.Println(armM.MachineID)
			ProductGenARM(armM)
		}
		fmt.Println("Hello Cuh")
		fmt.Println(TotalParts)
		time.Sleep(1000 * time.Millisecond)
	}
}
