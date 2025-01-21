package landmarks

import (
	"fmt"
	"io"
	"os"
	"strings"

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

	data = YamlComments(data)

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

func YamlComments(original []byte) []byte {
	originalLines := strings.Split(string(original), "\n")
	
	final := []string{"# Landmarks file\n\n"}
	final = append(final, fmt.Sprintf("%s # Each array element represents a landmark \n", originalLines[0]))
	final = append(final, fmt.Sprintf("# %s  - this will be used in commands", originalLines[1]))
	final = append(final, fmt.Sprintf("# %s  - absolute path to landmark", originalLines[2]))
	final = append(final, fmt.Sprintf("# %s  - landmark type - can be 'dir', 'file' or 'app'. Each type offers different operations \n", originalLines[3]))

	return []byte(strings.Join(final, "\n"))
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
