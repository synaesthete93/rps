package landmarks

import (
	"fmt"
	"io"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func Path() string {
	home, err := os.UserHomeDir()

	if err != nil {
		panic("Could not get landmarks file path: " + err.Error())
	}

	return home + "/.landmarks.yaml"
}

func InitLandmarksFile() {
	landmarks := InitLandmarks()

	data, err := yaml.Marshal(landmarks)
	if err != nil {
		panic("Could not initialize landmarks file: " + err.Error())
	}

	file, err := os.Create(Path())
	if err != nil {
		panic("Could not initialize landmarks file: " + err.Error())
	}

	defer file.Close()

	_, err = file.Write(data)

	if err != nil {
		panic("Could not initialize landmarks file: " + err.Error())
	}
}

func GetLandmarks() (*LandmarkFile, error) {
	file, err := os.Open(Path())
	if err != nil {
		if os.IsNotExist(err) {
			// Handle the case where the file does not exist
			return nil, fmt.Errorf("landmarks file does not exist: %w", err)
		}
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var landmarkFile LandmarkFile
	err = yaml.Unmarshal(data, &landmarkFile)
	if err != nil {
		return nil, err
	}

	return &landmarkFile, nil
}
