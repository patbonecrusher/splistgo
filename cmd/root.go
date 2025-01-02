/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"go.bug.st/serial/enumerator"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "splistgo",
	Short: "List serial port devices",
	Long: `Stupid dead simple app that allows one to print
	the path to a serial port device using it's serial number.
	
	Example:
		splistgo 1234567890 will print the path to this device
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.splistgo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
