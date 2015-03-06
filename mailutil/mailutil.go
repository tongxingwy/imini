package MailUtil

import(
  "github.com/tongxingwy/gomail"
  )

func SendEmail(){
  msg := gomail.NewMessage()
  msg.SetHeader("From", "alex@example.com")
  msg.SetHeader("To", "bob@example.com", "cora@example.com")
  msg.SetAddressHeader("Cc", "dan@example.com", "Dan")
  msg.SetHeader("Subject", "Hello!")
  msg.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")

  f, err := gomail.OpenFile("/home/Alex/lolcat.jpg")
  if err != nil {
      panic(err)
  }
  msg.Attach(f)

  // Send the email to Bob, Cora and Dan
  mailer := gomail.NewMailer("smtp.example.com", "user", "123456", 587)
  err := mailer.Send(msg); err != nil {
      panic(err)
  }
}
