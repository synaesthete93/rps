package landmarks

import (
	"os"
	"fmt"
)

type LandmarkFile struct {
	Landmarks []Landmark
}

type Landmark struct {
	Name *string
	Path *string
	Type *LandmarkType
}

type LandmarkType string

const DirType LandmarkType = "dir"
const FileType LandmarkType = "file"
const AppType LandmarkType = "app"

func InitLandmarks() *LandmarkFile {
	landmarks := make([]Landmark, 1)

	dirType := DirType
	userHomeDir, _ := os.UserHomeDir()

	landmarks[0] = Landmark{
		Name: func(s string) *string { return &s }("home"),
		Path: &userHomeDir,
		Type: &dirType,
	}

	return &LandmarkFile{
		Landmarks: landmarks,
	}
}

func FindLandmark(name string) (*Landmark, error) {
	lmks, _ := GetLandmarks()
	for _, landmark := range lmks.Landmarks {
		fmt.Println(*landmark.Name)
		if *landmark.Name == name {
			return &landmark, nil
		}
	}

	return nil, nil
}

//TODO: add validation for LandmarkType
