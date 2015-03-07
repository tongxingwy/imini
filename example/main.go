package main

import (
    "fmt"
    "github.com/tongxingwy/imini"
    "github.com/tongxingwy/imini/dbhelper"
    //"log"
)

func main() {
    data := make(map[string]interface{})
    config,err := imini.LoadConfig("../config/config.xml")
    //log.Println(config)
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
  }
  //log.Println("data:",data)
  if config.Tmpl.IsOn == "true"{
    imini.RenderHtml(config.Tmpl,data)
  }

  if config.Email.IsOn == "true"{
    imini.SendEmail(config.Email)
  }

}
