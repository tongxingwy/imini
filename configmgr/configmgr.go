package ConfigManager

import (
  "bytes"
  "github.com/tongxingwy/imini/dom4g"
  "io/ioutil"
  "log"
  )

var doc *dom4g.Element

func LoadConfig(fileUrl string){
  if doc != nil{
    log.Println("already load config file !!")
    return
  }
  buf,err := ioutil.ReadFile(fileUrl)
  if err != nil {
    panic(err.Error())
  }
  reader := bytes.NewReader(buf)
  doc,err = dom4g.LoadByStream(reader)
  if err != nil {
    panic(err.Error())
  }
}

func GetDbs()
