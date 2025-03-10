/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"os/exec"
	"fmt"
	"github.com/synaesthete93/rps/landmarks/pkg/landmarks"
	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goto",
	Short: "Go to a landmark",
	Long: `Jumps to a user-defined landmark`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		landmarkName:= args[0]
		if landmarkName == "" {
			cmd.Help()
			os.Exit(1)
		}
		
		landmark, err := landmarks.FindLandmark(landmarkName)

		if err != nil {
			panic("Could not find landmark: " + err.Error())
		}

		if landmark == nil {
			fmt.Println(fmt.Sprintf("There is no landmark named '%s'", landmarkName))
			fmt.Println(fmt.Sprintf("You can add it manually in %s", landmarks.Path()))
			fmt.Println("OR")
			fmt.Println("You can add it with the 'landmarks add' command")
			os.Exit(1)
		}
		targetDir := *landmark.Path

		command := exec.Command(os.Getenv("SHELL"))
		command.Dir = targetDir
		command.Stdin = os.Stdin
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		
		command.Run()
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
	rootCmd.Args = cobra.MinimumNArgs(1)
}


