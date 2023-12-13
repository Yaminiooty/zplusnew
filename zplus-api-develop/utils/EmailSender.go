package utils

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/go-mail/mail"
)

func SendEmailWithToolReport(sendToEmailId string, fullFilePath string, username string) error {
	smtpHost := os.Getenv("EMAIL_SMTP_HOST")
	smtpPort, _ := strconv.Atoi(os.Getenv("EMAIL_SMTP_PORT"))
	htmlContent, err := ioutil.ReadFile(os.Getenv("EMAIL_REPORT_TEMPLATE_PATH"))
	if err != nil {
		return err
	}
	formattedHTML := fmt.Sprintf(string(htmlContent))
	formattedHTML = strings.Replace(formattedHTML, "frontendurl", os.Getenv("FRONTEND_SERVER_URL")+"/email", 8)
	formattedHTML = strings.Replace(formattedHTML, "--user--", username, 1)
	m := mail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_FROM"))
	m.SetHeader("To", sendToEmailId)
	m.SetHeader("Subject", "Z+ Tool Execution Report")
	m.SetBody("text/html", formattedHTML)
	m.Attach(fullFilePath + ".html")
	m.Attach(fullFilePath + ".pdf")
	d := mail.NewDialer(smtpHost, smtpPort, os.Getenv("EMAIL_FROM"), os.Getenv("EMAIL_FROM_PASSWORD"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func SendEmail(userName string, sendToEmailId string, resetPasswordToken string) error {
	smtpHost := os.Getenv("EMAIL_SMTP_HOST")
	smtpPort, _ := strconv.Atoi(os.Getenv("EMAIL_SMTP_PORT"))
	htmlContent, err := ioutil.ReadFile(os.Getenv("RESET_PASSWORD_TEMPLATE_PATH"))
	if err != nil {
		return err
	}
	formattedHTML := fmt.Sprintf(string(htmlContent))
	formattedHTML = strings.Replace(formattedHTML, "frontendurl", os.Getenv("FRONTEND_SERVER_URL")+"/email", 9)
	formattedHTML = strings.Replace(formattedHTML, "--user--", userName, 1)
	resetPasswordLink := os.Getenv("FRONTEND_SERVER_URL") + "/reset-password" + "?token=" + resetPasswordToken
	formattedHTML = strings.Replace(formattedHTML, "resetpasswordlink", resetPasswordLink, 1)
	m := mail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_FROM"))
	m.SetHeader("To", sendToEmailId)
	m.SetHeader("Subject", "Z+ Account Password Reset")
	m.SetBody("text/html", formattedHTML)
	d := mail.NewDialer(smtpHost, smtpPort, os.Getenv("EMAIL_FROM"), os.Getenv("EMAIL_FROM_PASSWORD"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func SendEmailWithAccountVerificationCode(userName string, sendToEmailId string, verificationCode int) error {
	smtpHost := os.Getenv("EMAIL_SMTP_HOST")
	smtpPort, _ := strconv.Atoi(os.Getenv("EMAIL_SMTP_PORT"))
	htmlContent, err := ioutil.ReadFile(os.Getenv("EMAIL_TEMPLATE_PATH"))
	if err != nil {
		return err
	}
	formattedHTML := fmt.Sprintf(string(htmlContent))
	formattedHTML = strings.Replace(formattedHTML, "frontendurl", os.Getenv("FRONTEND_SERVER_URL")+"/email", 8)
	formattedHTML = strings.Replace(formattedHTML, "frontendloginpage", os.Getenv("FRONTEND_SERVER_URL"), 1)
	formattedHTML = strings.Replace(formattedHTML, "--user--", userName, 1)
	formattedHTML = strings.Replace(formattedHTML, "--code--", strconv.Itoa(verificationCode), 1)
	m := mail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_FROM"))
	m.SetHeader("To", sendToEmailId)
	m.SetHeader("Subject", "Z+ Account Verification: Welcome Aboard!")
	m.SetBody("text/html", formattedHTML)
	d := mail.NewDialer(smtpHost, smtpPort, os.Getenv("EMAIL_FROM"), os.Getenv("EMAIL_FROM_PASSWORD"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
