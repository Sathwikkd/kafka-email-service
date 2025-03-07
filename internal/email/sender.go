package email

import (
	"bytes"
	"html/template"
	"log"
	"net/smtp"
)

// SMTP Config (Update with your actual SMTP settings)
const (
	SMTPServer = "smtp.gmail.com"
	SMTPPort   = "587"
	SMTPUser   = "noreply.vithsutra@gmail.com"
	SMTPPass   = "vlco ctlo uuzm wqqv"
)

// SendEmail function to send emails using HTML templates
func SendEmail(to, subject, templateFile string, data map[string]string) error {
	// Parse the HTML template
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Printf("[ERROR] Failed to parse template %s: %v", templateFile, err)
		return err
	}

	// Replace placeholders with actual data
	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		log.Printf("[ERROR] Failed to execute template: %v", err)
		return err
	}

	// Set email headers
	message := "MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"From: " + SMTPUser + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n\r\n" +
		body.String()

	// Send the email
	auth := smtp.PlainAuth("", SMTPUser, SMTPPass, SMTPServer)
	err = smtp.SendMail(SMTPServer+":"+SMTPPort, auth, SMTPUser, []string{to}, []byte(message))
	if err != nil {
		log.Printf("[ERROR] Failed to send email: %v", err)
		return err
	}

	log.Printf("[INFO] Email sent successfully to %s", to)
	return nil
}
