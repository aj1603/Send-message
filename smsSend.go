package main

import (
	"bytes"
	"log"
	"os/exec"
	"time"
)

func main() {
	handle()
}

func handle() {
	otp := "1TL=0,67tmt onunden tolegsiz! Aksiýa aprelin 23-31 aralygy dowam edýär."

	phones := []string{
		"71193866", "71193866", "71193866", "71193866", "71193866", "71193866",
		"61581331", "61581331", "61581331", "61581331", "61581331", "61581331", "61581331",
	}

	chunkSize := 5

	for chunkIndex := 0; chunkIndex < len(phones); chunkIndex += chunkSize {
		end := chunkIndex + chunkSize
		if end > len(phones) {
			end = len(phones)
		}
		phoneChunk := phones[chunkIndex:end]

		for _, phone := range phoneChunk {
			phoneWithCode := "993" + phone
			command := "echo jora1603 | sudo -S gammu -c ~/.gammurc sendsms TEXT " + phoneWithCode + " -unicode -text \"" + otp + "\""

			cmd := exec.Command("bash", "-c", command)
			var out bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &out

			err := cmd.Run()
			output := out.String()

			if err != nil {
				log.Printf("Error sending SMS to %s: %s\n", phoneWithCode, output)
			} else {
				log.Printf("SMS sent successfully to %s: %s\n", phoneWithCode, output)
			}
		}

		// Sleep for 8 seconds after each chunk, but not after the last chunk
		if chunkIndex+chunkSize < len(phones) {
			time.Sleep(8 * time.Second)
		}
	}
}
