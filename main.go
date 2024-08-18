package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type MessageRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Text        string `json:"text" binding:"required"`
}

func main() {
	r := gin.Default()

	r.StaticFile("/", "./index.html")
	r.Static("/static", "./")

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/send-message", func(c *gin.Context) {
		var messageRequest MessageRequest
		if err := c.ShouldBindJSON(&messageRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		phoneWithCode := "993" + string(messageRequest.PhoneNumber)
		command := fmt.Sprintf("echo gerekli1603 | sudo -S gammu -c ~/.gammurc sendsms TEXT %s -unicode -len %s -text \"%s\"", phoneWithCode, getLengthOption(messageRequest.Text), messageRequest.Text)

		cmd := exec.Command("bash", "-c", command)
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &out

		err := cmd.Run()
		output := out.String()
		log.Println(phoneWithCode)

		if err != nil {
			if strings.Contains(output, "Failed to get SMSC number from phone") {
				log.Printf("Error sending SMS to %s: Failed to get SMSC number from phone\n", phoneWithCode)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get SMSC number from phone"})
			} else {
				log.Printf("Error sending SMS to %s: %s\n", phoneWithCode, output)
				c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("SMS sending failed: %s", err)})
			}
		} else {
			log.Printf("SMS sent successfully to %s: %s\n", phoneWithCode, output)
			c.JSON(http.StatusOK, gin.H{"message": "SMS sent successfully"})
		}
	})

	r.Run(":8081")
}

func getLengthOption(message string) string {
	messageLength := len(message)
	if messageLength <= 70 {
		return "70"
	}
	return "200"
}
