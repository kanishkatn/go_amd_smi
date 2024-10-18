package main

import "log"

func main() {
	if i := InitAMDSMI(); !i {
		log.Println("Failed to initialize AMD SMI")
		return
	}

	log.Println("AMD SMI initialized")

	n, err := NumAMDSMIMonitorDevices()
	if err != nil {
		log.Println("failed to get number of monitor devices:", err)
		return
	}

	log.Println("Number of monitor devices:", n)

	for i := 0; i < int(n); i++ {
		name, err := GetAMDSMIDeviceName(i)
		if err != nil {
			log.Println("failed to get device name:", err)
			return
		}

		log.Println("Device name:", name)
	}

	i, err := ShutdownAMDSMI()
	if err != nil {
		log.Println("failed to shutdown AMD SMI:", err)
		return
	}

	if !i {
		log.Println("Failed to shutdown AMD SMI")
		return
	}

	log.Println("AMD SMI shutdown")
}
