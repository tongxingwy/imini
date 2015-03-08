package imini

import (
  "encoding/xml"
  "fmt"
  "io/ioutil"
  )

type _Db struct{
  Key string `xml:"key,attr"`
  IsOn string `xml:"isOn,attr"`
  Url string `xml:"url"`
  Sql string `xml:"sql"`
}

type _Tmpl struct{
  IsOn string `xml:"isOn,attr"`
  FileIn string `xml:"filein"`
  FileOut string `xml:"fileout"`
}

type _Email struct{
  IsOn string `xml:"isOn,attr"`
  From string `xml:"from"`
  To string `xml:"to"`
  Cc string `xml:"cc"`
  //CcUser string `xml:"name.attr"`
  Sip string `xml:"sip"`
  Sport string `xml:"sport"`
  User string `xml:"user"`
  Pwd string `xml:"pwd"`
  Subject string `xml:"subject"`
  Content string `xml:"content"`
  Attach string `xml:"attach"`
}

type Config struct{
  Email _Email `xml:"email"`
  Tmpl _Tmpl `xml:"tmpl"`
  Dbs []_Db `xml:"dbs>db"`
}

func LoadConfig(fileUrl string) (*Config,error){
  var config *Config
  buf,err := ioutil.ReadFile(fileUrl)
  if err != nil {
    panic(err.Error())
  }
  err = xml.Unmarshal([]byte(buf), &config)
  if err != nil {
      fmt.Printf("error: %v", err)
      return nil,err
  }
  return config,nil
}
