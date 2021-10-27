package email

import (
	"fmt"
	"net/smtp"
)

// SmtpServer data to smtp server
type smtpServer struct {
	host string
	port string
}

// Address URI to smtp server
func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

// Do send email
func Do() {
	configuration := GetConfig()
	fmt.Printf("Username is %s and password is %s \n", configuration.Username, configuration.Password)

}

// Send email
func Send(body string, subject string) {
	config := GetConfig()
	from := config.Username
	password := config.Password

	to := []string{
		"dipesh.pradhan@outlook.com",
	}

	// Smtp server configuration
	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}

	//body := "This is a sample email."
	msg := "From: " + from + "\n" +
		"To: " + to[0] + "\n" +
		"Subject: " + subject + "\n\n" + body
	// Message
	message := []byte(msg)

	// Authentication
	auth := smtp.PlainAuth("", from, password, smtpServer.host)

	// Sending email
	err := smtp.SendMail(smtpServer.Address(), auth, from, to, message)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Email sent")
}
