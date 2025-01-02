/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"go.bug.st/serial/enumerator"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List serial port devices",
	Long:  `List your connected serial port devices.`,
	Run: func(cmd *cobra.Command, args []string) {
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
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
