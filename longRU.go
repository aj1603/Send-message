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
	otp := "GEREKLI GLOBAL 1 TL = 0,67 TMT ÖŇUNDEN TÖLEGSIZ! Aksiýa 23-nji aprelden 31-nji aprel aralygynda dowam edýär." +
		"1 TL = 0,67 ТМТ БЕЗ ПРЕДОПЛАТЫ! Акция продлится с 23 по 31 апреля. ✈️ Доставка самолетом: 1 кг/98 тмт (7-15 дней) 📦 Доставка контейнером: 1 кг/20 тмт (30-45 дней)"

	phones := []string{
		"71193866", "71193866",
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
			// Use single quotes to enclose the entire command to avoid issues with special characters
			command := `echo jora1603 | sudo -S gammu -c ~/.gammurc sendsms TEXT ` + phoneWithCode + ` -unicode -text "` + otp + `"`

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

// getLengthOption returns the appropriate length option for gammu based on the message length
func getLengthOption(message string) string {
	messageLength := len(message)
	if messageLength <= 70 {
		return "70"
	}
	return "200"
}
