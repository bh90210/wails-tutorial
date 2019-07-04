package wailscomm

import (
	"fmt"
	sc "simpleclient/api/serverCommunication"
)

type ServiceChooser struct {
	Switch string
}

func NewserviceChooser() *ServiceChooser {
	result := &ServiceChooser{
		Switch: "0",
	}
	return result
}

// Choose .
func (sw *ServiceChooser) Choose(service string) {
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
	fmt.Sprintf("My name is now '%s'", service)
}
