package main

import (
    "fmt"
    "github.com/tongxingwy/imini"
    "github.com/tongxingwy/imini/dbhelper"
    "log"
)

func main() {
    data := make(map[string]interface{})
    isReadDb := false
    configFile := "./config/config.xml"
    config,err := imini.LoadConfig(configFile)
    log.Println("load config file ",configFile," successfully...")

    if err != nil{
      panic(err.Error())
    }
    for _,db := range config.Dbs {
      if db.IsOn == "false"{
        continue
      }
      // Open database connection
      conn,err := DriverManager.GetConnection(db.Url)
      if err != nil {
        fmt.Println("database initialize error : ",err.Error());
        panic(err.Error())
      }
      result := conn.Query(db.Sql)
      data[db.Key] = result
      conn.Close()
      isReadDb = true
  }

  if isReadDb {
    log.Println("fectch data from database complete...")
  }

  if config.Tmpl.IsOn == "true"{
    imini.RenderHtml(config.Tmpl,data)
    log.Println("parse template over,output file is: ",config.Tmpl.FileOut)
  }

  if config.Email.IsOn == "true"{
    imini.SendEmail(config.Email)
    log.Println("send email successfully...")
  }
  log.Println("over-----------------------------")
}
