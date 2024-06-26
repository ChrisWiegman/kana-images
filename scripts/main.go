package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Image struct {
	Src string
	Tag string
}

var imageBasePath = "../images/"
var dockerOrg = "chriswiegman/kana"
var images = []Image{}
var Stdout io.Writer = os.Stdout
var Stderr io.Writer = os.Stderr

func main() {
	images := getImages(imageBasePath, images)

	for _, image := range images {
		tag := fmt.Sprintf("%s:%s", dockerOrg, image.Tag)

		dockerCommands := [][]string{
			{"buildx", "create", "--use"},
			{"buildx", "build", "--push", "--platform", "linux/amd64,linux/386,linux/arm/v6,linux/arm/v7,linux/arm64", "-t", tag, filepath.Join(imageBasePath, image.Src)},
		}

		for _, dockerCommand := range dockerCommands {
			cmd := exec.Command("docker", dockerCommand...)

			cmd.Stdout = Stdout
			cmd.Stderr = Stderr
			err := cmd.Run()
			if err != nil {
				panic(err)
			}
		}
	}
}

func getImages(directory string, images []Image) []Image {
	items, _ := os.ReadDir(directory)

	for _, item := range items {
		if item.IsDir() {
			images = getImages(filepath.Join(directory, item.Name()), images)
		} else {
			if item.Name() == "Dockerfile" {
				tag := strings.Replace(
					strings.Replace(directory, imageBasePath, "", 1),
					"/", "-", -1)

				image := Image{directory, tag}
				images = append(images, image)
			}
		}
	}

	return images
}
