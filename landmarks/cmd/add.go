/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/synaesthete93/rps/landmarks/pkg/landmarks"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a landmark",
	Long:  `Adds a landmark to the list of landmarks.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Error: must pass exactly one argument representing the landmark path")
			return
		}

		landmarkPath := args[0]

		processedLandmarkPath, err := processLandmarkPath(landmarkPath)

		if err != nil {
			fmt.Printf("Error processing landmark path: %v\n", err)
			return
		}

		var landmarkName string
		for {
			landmarkName, err = promptForLandmarkName()
			if err == nil {
				break
			}
			fmt.Printf("Error: %v\n", err)
		}

		landmarkType, err := decideType(processedLandmarkPath)

		if err != nil {
			fmt.Printf("Error deciding landmark type: %v\n", err)
			return
		}

		if !confirmLandmark(landmarkName, processedLandmarkPath, landmarkType) {
			fmt.Println("Landmark addition cancelled.")
			return
		}

		newLandmark := landmarks.Landmark{
			Name: &landmarkName,
			Type: &landmarkType,
			Path: &processedLandmarkPath,
		}

		landmarks.AddLandmark(newLandmark)

		fmt.Println("Landmark added successfully.")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func promptForLandmarkName() (string, error) {
	prompt := promptui.Prompt{
		Label: "Enter the landmark name",
	}

	result, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("prompt failed: %w", err)
	}

	if result == "" {
		return "", fmt.Errorf("landmark name cannot be empty")
	}

	existing, _ := landmarks.FindLandmark(result)

	if existing != nil {
		return "", fmt.Errorf("landmark with name '%s' already exists", result)
	}

	return result, nil
}

func processLandmarkPath(path string) (string, error) {
	if path == "" {
		return "", fmt.Errorf("path cannot be empty")
	}

	if filepath.IsAbs(path) {
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}

	if path == "." {
		wd, err := os.Getwd()
		if err != nil {
			return "", fmt.Errorf("could not get current directory: %v", err)
		}
		return wd, nil
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("could not get absolute path: %v", err)
	}

	if _, err := os.Stat(absPath); err == nil {
		return absPath, nil
	} else if os.IsNotExist(err) {
		return "", fmt.Errorf("path does not exist: %v", err)
	} else {
		return "", fmt.Errorf("could not stat path: %v", err)
	}
}

func decideType(path string) (landmarks.LandmarkType, error) {
	info, err := os.Stat(path)
	if err != nil {
		return "", fmt.Errorf("could not stat path: %v", err)
	}

	if info.Mode().IsRegular() {
		fmt.Printf("Path is a file. Choosing 'file' type.\n")
		return landmarks.FileType, nil
	}

	if info.Mode().IsDir() {
		typeChoice, err := chooseDirOrApp()

		if err != nil {
			return "", fmt.Errorf("could not choose dir or app: %v", err)
		}

		if typeChoice == "directory" {
			return landmarks.DirType, nil
		}

		if typeChoice == "app" {
			return landmarks.AppType, nil
		}
	}

	return "", fmt.Errorf("could not decide landmark type")
}

func chooseDirOrApp() (string, error) {
	prompt := promptui.Select{
		Label: "Should this be treated as an app or a regular directory?",
		Items: []string{"directory", "app"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("prompt failed: %w", err)
	}

	return result, nil
}

func confirmLandmark(landmarkName string, processedLandmarkPath string, landmarkType landmarks.LandmarkType) bool {
	fmt.Println()
	fmt.Printf("Landmark name: %s\n", landmarkName)
	fmt.Printf("Landmark path: %s\n", processedLandmarkPath)
	fmt.Printf("Landmark type: %s\n", landmarkType)
	fmt.Println()

	prompt := promptui.Prompt{
		Label:     "Do you want to add this landmark",
		IsConfirm: true,
	}

	_, err := prompt.Run()
	if err != nil {
		fmt.Println("Landmark addition cancelled.")
		return false
	}

	return true
}
