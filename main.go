package main

import (
	"fmt"
	"log"
	"os/exec"
)

func send(phone, otp string) {
	phoneWithCode := fmt.Sprintf("993%s", phone)
	command := fmt.Sprintf(
		"echo jora1603 | sudo -S gammu -c ~/.gammurc sendsms TEXT %s -unicode -text \"%s\"",
		phoneWithCode, otp,
	)

	cmd := exec.Command("/usr/bin/fish", "-c", command)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
		log.Printf("Output: %s\n", string(output))
	} else {
		log.Printf("Success: %s\n", string(output))
	}
}

// func main() {
// 	send("71193866", "Приведи друга и получи скидку на заказы из-за рубежа! Условия просты приводи друзей,пусть они сделают заказ, укажи номер и имя друга которого привел при заказе и мы сделаем скидку! +99364021049")
// }

func main() {
	send("71193866", "1 TL = 0, 66 тмт БЕЗ ПРЕДОПЛАТЫ - ЗАКАЗЫ ИЗ ТУРЦИИ! АКЦИЯ ПРОДЛИТСЯ С 23-31 МАЯ Доставка самолетом: 1 кг/98 тмт (7-15 дней) 📦 Доставка контейнером: 1 кг/20 тмт (30-45 дней) 📞Для заказа свяжитесь по номеру 861-628-042 или напишите нам в Direct")
}
