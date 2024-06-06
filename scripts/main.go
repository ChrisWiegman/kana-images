package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

type Image struct {
	Src string
	Tag string
}

var imageBasePath = "../images/"
var dockerOrg = "chriswiegman/kana"

var images = []Image{
	{
		Src: "wordpress/cli/php8.1",
		Tag: "cli-php8.1",
	},
	{
		Src: "wordpress/cli/php8.2",
		Tag: "cli-php8.2",
	},
	{
		Src: "wordpress/cli/php8.3",
		Tag: "cli-php8.3",
	}}

func main() {
	for _, image := range images {
		tag := fmt.Sprintf("%s:%s", dockerOrg, image.Tag)

		dockerCommands := [][]string{
			{"build", "-t", tag, filepath.Join(imageBasePath, image.Src)},
			{"push", tag},
		}

		// Stdout is the io.Writer to which executed commands write standard output.
		var Stdout io.Writer = os.Stdout

		// Stderr is the io.Writer to which executed commands write standard error.
		var Stderr io.Writer = os.Stderr

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
