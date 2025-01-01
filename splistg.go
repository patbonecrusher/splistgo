package main

import (
	"fmt"
	"log"
	"os"

	"go.bug.st/serial/enumerator"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		listPorts()
	} else {
		findPort(args[0])
	}
}

func listPorts() {
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		return
	}
	for _, port := range ports {
		fmt.Printf("Port: %s\n", port.Name)
		if port.Product != "" {
			fmt.Printf("   Product Name: %s\n", port.Product)
		}
		if port.IsUSB {
			fmt.Printf("   USB ID      : %s:%s\n", port.VID, port.PID)
			fmt.Printf("   USB serial  : %s\n", port.SerialNumber)
		}
		fmt.Println()
	}
}

func findPort(serialNumber string) {
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		log.Fatal(err)
	}
	for _, port := range ports {
		if port.SerialNumber == serialNumber {
			fmt.Printf("%s", port.Name)
			return
		}
	}
	// fmt.Printf("Port with serial number %s not found\n", serialNumber)
}
