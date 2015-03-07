package imini

import (
  "text/template"
  "os"
  )

func RenderHtml(tmpl _Tmpl,data interface{}){
  t, err := template.ParseFiles(tmpl.FileIn)
  if err != nil {
    panic(err.Error())
  }
  filep, err := os.OpenFile(tmpl.FileOut, os.O_RDWR|os.O_TRUNC|os.O_CREATE, os.ModePerm)
  if err != nil {
      panic(err)
  }
  defer filep.Close()
  err = t.Execute(filep, data)
  if err != nil {
    panic(err.Error())
  }
}
