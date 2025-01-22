/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/synaesthete93/rps/landmarks/pkg/landmarks"
)

// pathCmd represents the path command
var pathCmd = &cobra.Command{
	Use:   "path [landmark]",
	Short: "Gets the path to the landmarks file or a specific landmark",
	Long: `Gets the path to the landmarks file or a specific landmark
	If no argument is provided, returns the path to the landmarks file. 
	If one argument is provided, returns the path to the landmark with that name.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(landmarks.Path())
			return
		}

		landmarkName := args[0]

		landmark, err := landmarks.FindLandmark(landmarkName)

		if err != nil {
			fmt.Printf("Error getting landmark: %v\n", err)
			fmt.Println()
			return
		}

		if landmark == nil {
			fmt.Printf("Landmark %s not found", landmarkName)
			fmt.Println()
			return
		}

		fmt.Println(*landmark.Path)
	},
}

func init() {
	rootCmd.AddCommand(pathCmd)
}
