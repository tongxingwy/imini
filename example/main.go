package main

import (
    "fmt"
    //"github.com/tongxingwy/imini"
    "github.com/tongxingwy/imini/dbhelper"
    //"github.com/tongxingwy/imini/mailutil"
    //"log"
)

func main() {
    // Open database connection
    conn,err := DriverManager.GetConnection("aiuap:aiuap@tcp(10.1.198.135:3306)/ai4a")
    if err != nil {
      fmt.Println("database initialize error : ",err.Error());
      panic(err.Error())
    }
    conn.Query("SELECT count(*) '4Amini活跃度' FROM uap_main_acct t WHERE t.last_login_time BETWEEN CURDATE() AND NOW() UNION SELECT count(*) FROM uap_main_acct t UNION SELECT concat(round((SELECT count(*) FROM uap_main_acct t WHERE t.last_login_time BETWEEN CURDATE() AND NOW()) / ((SELECT count(*) FROM uap_main_acct t)) * 100,2), '%')")
    conn.Close()
    
}
