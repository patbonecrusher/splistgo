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

// pathCmd represents the path command
var pathCmd = &cobra.Command{
	Use:   "path",
	Short: "Return the path to a serial port device using it's serial number",
	Long: `Useful on macOS or Linux to find device path when
multiple devices are connected. For example:

ls ` + "`splistgo path 74:4D:BD:75:22:F4`" + `

will print the path to the device with serial number 74:4D:BD:75:22:F4.
"/dev/cu.usbmodem112301" or "/dev/ttyUSB0" or similar.
`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("Serial number is required")
		}

		ports, err := enumerator.GetDetailedPortsList()
		if err != nil {
			log.Fatal(err)
		}
		for _, port := range ports {
			if port.SerialNumber == args[0] {
				fmt.Printf("%s", port.Name)
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(pathCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pathCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pathCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
