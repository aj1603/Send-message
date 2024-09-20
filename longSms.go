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

	otp := "😍Телевизор LUCK диагональю 50 всего за 373 тмт в месяц 🎁📺  \n" +
		"Условия акции просты: \n" +
		"1️⃣Подойдите в офис Gerekli с паспортом (мы лишь снимаем копию паспорта, после чего документ возвращается обратно к вам) 2️⃣Оплатите первоначальный взнос в размере 373 тмт и ТЕЛЕВИЗОР ВАШ❗️😃 \n" +
		"По дополнительным вопросам обращайтесь по номеру ☎️27-10-15☎️ \n" +
		"😍Luck telewizory, 50-lik diagonally, bary ýogy 373 manat her aý tölegli🎁📺 \n" +
		"Aksiýanyň şertleri ýönekeý: \n" +
		"1️⃣Pasport bilen Gerekli ofisine geliň (biz diňe pasport nusgasyny alyp, resminamany yzyna gaýtarýarys) \n" +
		"2️⃣Oňünden geçirilýän 373 manady töläň we telewizor siziňki! \n" +
		"Goşmaça soraglar boýunça ☎️27-10-15☎️ telefon belgisine jaň ediñ"

	phones := []string{"71193866"}
	chunkSize := 5

	for chunkIndex := 0; chunkIndex < len(phones); chunkIndex += chunkSize {
		end := chunkIndex + chunkSize
		if end > len(phones) {
			end = len(phones)
		}
		phoneChunk := phones[chunkIndex:end]

		for _, phone := range phoneChunk {
			phoneWithCode := "993" + phone
			command := "echo gerekli1603 | sudo -S gammu -c ~/.gammurc sendsms TEXT " + phoneWithCode + " -unicode -len " + getLengthOption(otp) + " -text \"" + otp + "\""

			cmd := exec.Command("bash", "-c", command)
			var out bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &out

			err := cmd.Run()
			output := out.String()
			log.Println(phone)

			if err != nil {
				log.Printf("Error sending SMS to %s: %s\n", phoneWithCode, output)
			} else {
				log.Printf("SMS sent successfully to %s: %s\n", phoneWithCode, output)
			}
		}

		if chunkIndex+chunkSize < len(phones) {
			time.Sleep(15 * time.Second)
		}
	}
}

func getLengthOption(message string) string {
	messageLength := len(message)
	if messageLength <= 70 {
		return "70"
	}
	return "200"
}
