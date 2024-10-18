package main

/*
#cgo CFLAGS: -Wall -I./goamdsmi_shim/amdsmi -I./goamdsmi_shim/goamdsmi -DENABLE_DEBUG_LEVEL=2 -DAMDSMI_BUILD
#include <stdint.h>
#include "goamdsmi_shim.h"
#include "amdsmi.h"

#include "amdsmi_go_shim.c"
*/
import "C"
import "fmt"

var initialized bool

// InitAMDSMI initializes the AMD SMI library.
func InitAMDSMI() bool {
	if initialized {
		return true
	}
	initialized = C.go_shim_amdsmigpu_init() == C.GOAMDSMI_STATUS_SUCCESS
	return initialized
}

// ShutdownAMDSMI shuts down the AMD SMI library.
func ShutdownAMDSMI() (bool, error) {
	if !initialized {
		return false, fmt.Errorf("AMD SMI library not initialized")
	}
	return C.go_shim_amdsmigpu_shutdown() == C.GOAMDSMI_STATUS_SUCCESS, nil
}

// NumAMDSMIMonitorDevices returns the number of monitor devices.
func NumAMDSMIMonitorDevices() (uint, error) {
	if !initialized {
		return 0, fmt.Errorf("AMD SMI library not initialized")
	}
	var num C.uint32_t
	status := C.go_shim_amdsmigpu_num_monitor_devices(&num)
	if status != C.GOAMDSMI_STATUS_SUCCESS {
		return 0, fmt.Errorf("failed to get number of monitor devices")
	}
	return uint(num), nil
}

// GetAMDSMIDeviceName returns the name of the device at the given index.
func GetAMDSMIDeviceName(i int) (string, error) {
	if !initialized {
		return "", fmt.Errorf("AMD SMI library not initialized")
	}

	var cStr *C.char
	status := C.go_shim_amdsmigpu_dev_name_get(C.uint(i), &cStr)
	if status != C.GOAMDSMI_STATUS_SUCCESS {
		return "", fmt.Errorf("failed to get device name, status: %d", status)
	}
	if cStr == nil {
		return "", fmt.Errorf("device name not found")
	}

	deviceName := C.GoString(cStr)
	return deviceName, nil
}
