/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/synaesthete93/rps/landmarks/pkg/landmarks"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes the landmarks file",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		filePath := landmarks.Path()
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			landmarks.InitLandmarksFile(false)
			return
		} else if err != nil {
			fmt.Printf("Error checking file: %v\n", err)
			return
		}

		if force, _ := cmd.Flags().GetBool("force"); force {
			landmarks.InitLandmarksFile(true)
		} else {
			fmt.Printf("Landmarks file already exists at %s.\n", filePath)
			fmt.Println("If you want to overwrite the file, run the command with the --force flag.")
			fmt.Println("WARNING: This will overwrite all existing landmarks.")

		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().BoolP("force", "f", false, "Force initialization of the landmarks file. This will overwrite all existing landmarks.")
}
