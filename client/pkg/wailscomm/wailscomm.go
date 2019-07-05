package wailscomm

import (
	"fmt"
	sc "simpleclient/api/serverCommunication"
)

// ServiceChooser .
type ServiceChooser struct {
	Switch string
}

// Choose .
func (sw *ServiceChooser) Choose(service string) string {
	sw.Switch = service
	//scanner := bufio.NewScanner(os.Stdin)
	//for {
	//  if scanner.Scan() {
	switch sw.Switch {
	case "1":
		sc.PausePrevRout()
		sc.CpuPlay <- struct{}{}
		sc.LastCall = "1"
	case "2":
		sc.PausePrevRout()
		sc.DiskPlay <- struct{}{}
		sc.LastCall = "2"
	case "3":
		sc.PausePrevRout()
		sc.LoadPlay <- struct{}{}
		sc.LastCall = "3"
	case "4":
		sc.PausePrevRout()
		sc.MemPlay <- struct{}{}
		sc.LastCall = "4"
	}
	//}
	//}
	return fmt.Sprintf("Service '%s'", service)
}
