package main

import (
	"email-notif/notifications"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendEmailHandler(c *gin.Context) {
	var emailDetails notifications.EmailCreds

	err := c.ShouldBindJSON(&emailDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request body"})
		return
	}

	if err := notifications.SendEmail(emailDetails); err != nil {
		log.Printf("Error sending email: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully"})

}
func main() {
	router := gin.Default()

	router.POST("/send/emails", SendEmailHandler)
	addr := "localhost:8080"
	log.Printf("The email server is running on port %s", addr)
	router.Run(addr)
}
