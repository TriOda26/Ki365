package main

import (
	// "github.com/ki365/ki365/server/abatement"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
)

var s = []string{
	"https://github.com/Ki365/Hardware.git",
	"https://github.com/Ki365/OLINUXINO.git",
	"https://github.com/Ki365/POV-Display.git",
	"https://github.com/Ki365/SPACEDOS01.git",
	"https://github.com/Ki365/ext-con-breakout-board.git",
	"https://github.com/Ki365/nuco-v.git",
	"https://github.com/Ki365/upsat-comms-hardware.git",
}

// Generates initial example projects from examples/source to examples/build
func main() {
	// abatement.GenerateExamples("./examples/source/", "./examples/build/", true)
	for i, v := range s {
		path := filepath.Join("examples/build/", fileNameWithoutExtension(v))
		log.Printf("Cloning repo (%d of %d) from link: %s to path %s\n", i+1, len(s), v, path)
		_, err := git.PlainClone(path, true, &git.CloneOptions{
			URL:      v,
			Progress: os.Stdout,
		})

		if errors.Is(err, git.ErrRepositoryAlreadyExists) {
			log.Println("Repository already exists, skipping...")
			continue
		} else if err != nil {
			log.Fatal(fmt.Errorf("Error in cloning examples: %s", err))
		}
	}

	log.Println("Done cloning repositories.")
}

// https://github.com/Ki365/example.git -> example
func fileNameWithoutExtension(fileName string) string {
	return filepath.Base(strings.TrimSuffix(fileName, filepath.Ext(fileName)))
}
