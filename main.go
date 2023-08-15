package main

import (
	"crypto/tls"
	"fmt"
	gomail "gopkg.in/gomail.v2"
	"flag"
	"os"
	"strings"
	"bufio"
)

var (
  attachmentFlag = flag.String("a", "", "attach a file to the email")
  toFlag = flag.String("to","", "set send destination (required)")
  smtpFlag = flag.String("smtp", "smtp.gmail.com", "set custom smtp server, default is google")
)

func main() {
  flag.Parse()

  var username string
  var password string

  if len(flag.Args()) < 2 {
    fmt.Fprintln(os.Stderr, "Error, missing username or password nofity <username> <password> -to <desination>")
    return 
  } else {
    username = flag.Args()[0]
    password = flag.Args()[1]

    if strings.Contains(password, "-") {
      fmt.Fprintln(os.Stderr, "Error, missing username or password nofity -to <desination> <username> <password>")
    }
  }
  if toFlag == nil {
    fmt.Fprintln(os.Stderr, "Error, missing to destination, please use format nofity -to <desination> <username> <password> ")
  }

  msgString := readStdin()
  sendMessage(username, password, *toFlag, *attachmentFlag, msgString, *smtpFlag)
}

func readStdin() string {

  scanner := bufio.NewScanner(os.Stdin)
  
  isTerminal := isInputFromTerminal()
  var inputString strings.Builder
  for scanner.Scan() {
    inputString.WriteString(scanner.Text())

    if isTerminal {
      return inputString.String()
    }
  }

  return inputString.String()
}

func isInputFromTerminal() bool {
  stat, _ := os.Stdin.Stat()
  return (stat.Mode() & os.ModeCharDevice) != 0
}

func sendMessage(username string, password string, to string, attachment string, msg string, smtp string) {
  client := gomail.NewMessage()

  client.SetHeader("From", username)

  client.SetHeader("To", to)

  client.SetBody("text/plain", msg)

  dialer := gomail.NewDialer(smtp, 587, username, password)


  dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

  if err := dialer.DialAndSend(client); err != nil {
      fmt.Println(err)
      panic(err)
    }
}
