package main

import (
  "fmt"
  apns "github.com/anachronistic/apns"
)

func main() {
  var tokens = make([]string, 2)
  tokens[0] = "277d3346edde9d529b085eca27f6e98b5a542e3f7fdfee0a10451573f1fc7fbf"
  tokens[1] = "f28f7996b9d96316ee0d4ed2633665cdad04d86339ddeecdee412ed11a8911f5"

  for _, deviceToken := range tokens {
    payload := apns.NewPayload()
    payload.Alert = "今天几点吃饭，我是消息"
    payload.Badge = 42
    payload.Sound = "bingbong.aiff"

    pn := apns.NewPushNotification()
    pn.DeviceToken = deviceToken
    pn.AddPayload(payload)

    client := apns.NewClient("gateway.push.apple.com:2195", "apns-product-cert.pem", "apns-product-key-noenc.pem")
    resp := client.Send(pn)

    fmt.Println("apple return :", resp)

    alert, _ := pn.PayloadString()
    fmt.Println("  Alert:", alert)
    fmt.Println("Success:", resp.Success)
    fmt.Println("  Error:", resp.Error)
  }

}
