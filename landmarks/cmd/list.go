/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/synaesthete93/rps/landmarks/pkg/landmarks"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the landmarks available",
	Long: `Lists all the landmarks available, optionally along with their types.`,
	Run: func(cmd *cobra.Command, args []string) {
		 list, _ := landmarks.GetLandmarks()

		 for _, landmark := range list.Landmarks {
			fmt.Println(format(landmark, cmd.Flag("verbose").Value.String() == "true"))
			fmt.Println()
		 }
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("verbose", "v", false, "Display detailed information")
}

func format(landmark landmarks.Landmark, verbose bool) string {
	ret := fmt.Sprintf(" * %s", *landmark.Name)
	if verbose {
		ret += fmt.Sprintf(" (%s)", *landmark.Type)
	}

	ret += ":\n"

	ret += fmt.Sprintf("     Path: %s", *landmark.Path)

	return ret;
}