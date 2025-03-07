package models

// EmailRequest represents the email data structure
type EmailRequest struct {
	MailID   string `json:"mailid"`
	OTP      string `json:"otp,omitempty"`       // Only for OTP emails
	Username string `json:"username,omitempty"`  // Only for Welcome emails
	Platform string `json:"platform,omitempty"`  // Only for Welcome emails
}
