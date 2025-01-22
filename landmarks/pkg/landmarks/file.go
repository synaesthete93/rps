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

func InitLandmarksFile(overwrite bool) {
	landmarks := InitLandmarks()

	data, err := yaml.Marshal(landmarks)
	if err != nil {
		panic("Could not initialize landmarks file: " + err.Error())
	}

	data = YamlComments(data)

	writeToFile(data, overwrite)
}

func AddLandmark(newLandmark Landmark) {
	landmarksFile, err := GetLandmarks()
	if err != nil {
		panic("Could not get landmarks: " + err.Error())
	}

	landmarksFile.Landmarks = append(landmarksFile.Landmarks, newLandmark)

	SaveLandmarks(landmarksFile)
}

func RemoveLandmark(name string) {
	landmarksFile, err := GetLandmarks()
	if err != nil {
		panic("Could not get landmarks: " + err.Error())
	}

	var updatedLandmarks []Landmark
	for _, landmark := range landmarksFile.Landmarks {
		if *landmark.Name != name {
			updatedLandmarks = append(updatedLandmarks, landmark)
		}
	}

	landmarksFile.Landmarks = updatedLandmarks

	SaveLandmarks(landmarksFile)
}

func SaveLandmarks(landmarks *LandmarkFile) {
	data, err := yaml.Marshal(landmarks)
	if err != nil {
		panic("Could not save landmarks file: " + err.Error())
	}

	writeToFile(data, true)
}

func writeToFile(data []byte, overwrite bool) {
	if !overwrite {
		if _, err := os.Stat(Path()); err == nil {
			panic("Landmarks file already exists and overwrite is set to false")
		}
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
			InitLandmarksFile(false)
			return nil, fmt.Errorf("landmarks file does not exist. Creating now")
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
