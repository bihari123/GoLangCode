package main

import (
	"log"
	"os/exec"
)

func main() {
	s := "twelve thousand three hundred thirty-three"
	cmd := exec.Command("espeak", s)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
