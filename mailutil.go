package imini

import(
  "github.com/tongxingwy/gomail"
  "github.com/tongxingwy/imini/dateutil"
  "strconv"
  "regexp"
  "strings"
  "log"
  )

func SendEmail(email _Email){
  msg := gomail.NewMessage()
  msg.SetHeader("From", email.From)
  msg.SetHeader("To", email.To)
  if email.Cc != "" {
    msg.SetAddressHeader("Cc", email.Cc, email.CcUser)
  }
  msg.SetHeader("Subject", parseDataLayout(email.Subject))
  if email.Content!= "" {
    msg.SetBody("text/html", email.Content)
  }
  if email.Attach != "" {
    f, err := gomail.OpenFile(email.Attach)
    if err != nil {
        panic(err)
    }
    msg.Attach(f)
  }
  var sport int
  if port,err := strconv.ParseInt(email.Sport,10,32);err != nil{
    sport = 25
  }else{
    sport = int(port)
  }
  // Send the email to Bob, Cora and Dan
  mailer := gomail.NewMailer(email.Sip, email.User, email.Pwd, sport)
  if err := mailer.Send(msg); err != nil {
      panic(err)
  }
}

func parseDataLayout(subject string) string{
  re := regexp.MustCompile("\\{\\{([^\\{\\{]+)\\}\\}")
  groups := re.FindAllStringSubmatch(subject,-1)
  for _,pairs := range groups{
    log.Println(pairs,DateUtil.FormatNow(pairs[1]))
    subject = strings.Replace(subject, pairs[0], DateUtil.FormatNow(pairs[1]), -1)
  }
  return subject
}
