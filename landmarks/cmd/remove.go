/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/synaesthete93/rps/landmarks/pkg/landmarks"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove [landmark]",
	Short: "Removes a landmark",
	Long:  `Removes a landmark from the list of landmarks.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
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

		if confirmRemoval(landmark) {
			landmarks.RemoveLandmark(landmarkName)
			fmt.Printf("Landmark %s removed\n", landmarkName)
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func confirmRemoval(landmark *landmarks.Landmark) bool {
	fmt.Printf("Found landmark:\n")
	fmt.Printf("	Name: %s\n", *landmark.Name)
	fmt.Printf("	Type: %s\n", *landmark.Type)
	fmt.Printf("	Location: %s\n", *landmark.Path)
	fmt.Println()

	prompt := promptui.Prompt{
		Label:     fmt.Sprintf("Are you sure you want to remove the landmark %s", *landmark.Name),
		IsConfirm: true,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return false
	}

	return result == "y" || result == "Y"
}
