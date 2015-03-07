package main

import (
  "encoding/xml"
  "fmt"
  "github.com/tongxingwy/imini"
  "io/ioutil"
  "log"
  )

func main(){
  var config imini.Config
  buf,err := ioutil.ReadFile("../config/config.xml")
  if err != nil {
    panic(err.Error())
  }
  err = xml.Unmarshal([]byte(buf), &config)
  if err != nil {
      fmt.Printf("error: %v", err)
      return
  }
  log.Println(config)
}
