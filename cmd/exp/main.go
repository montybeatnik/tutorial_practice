package main

import (
	"bytes"
	"fmt"
)

func main() {
	d1 := deviceConnecter{
		IP: "1.1.1.1",
	}
	d2 := deviceConnecter{
		IP: "2.2.2.2",
	}
	var devs []DeviceConnector
	devs = append(devs, &d1, &d2)

	connectThemAll(devs)
}

type DeviceConnector interface {
	Connect() error
	RunCommand(c string) bytes.Buffer
}

type deviceConnecter struct {
	IP string
}

func (d *deviceConnecter) Connect() error {
	fmt.Printf("connected to %v\n", d.IP)
	return nil
}

func (d *deviceConnecter) RunCommand(c string) bytes.Buffer {
	var b bytes.Buffer
	return b
}

func connectThemAll(connectors []DeviceConnector) {
	for _, d := range connectors {
		fmt.Println(d.Connect())
	}
}
