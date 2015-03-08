package imini

import (
  "encoding/xml"
  "io/ioutil"
  "testing"
  )

func TestConfigParse(t *testing.T){
  var config Config
  buf,err := ioutil.ReadFile("./config/config.xml")
  if err != nil {
    t.Error(err)
  }
  err = xml.Unmarshal([]byte(buf), &config)
  if err != nil {
      t.Error(err)
  }
  t.Log(config)
}
