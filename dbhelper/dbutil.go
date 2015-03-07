package DriverManager

import(
  "database/sql"
    "fmt"
    _ "github.com/tongxingwy/mysql"
    //"strings"
    //"strconv"
  )


type Connection struct {
    db *sql.DB
}

func GetConnection(url string)(*Connection,error){
  conn := new(Connection)
  db,err := sql.Open("mysql",url)
  if err!=nil {
        fmt.Println("database initialize error : ",err.Error());
        return nil,err;
  }
  conn.db = db;
  return conn,nil;
}

func (conn *Connection)Update(){
    if conn.db==nil {
        return;
    }
    stmt,err := conn.db.Prepare("update test set name=?,age=? where age=?");
    if err!=nil {
        fmt.Println(err.Error());
        return;
    }
    defer stmt.Close();
    if result,err := stmt.Exec("周七",40,25);err==nil {
        if c,err := result.RowsAffected();err==nil {
            fmt.Println("update count : ",c);
        }
    }
}


func (conn *Connection)Query(sqlStr string) map[string]interface{}{
    if conn.db==nil {
      panic("please open database connection first!!")
    }
    // Execute the query
    rows, err := conn.db.Query(sqlStr)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    // Get column names
    columns, err := rows.Columns()
    //types := make([]int, len(columns))

    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    var data []interface{}
    // Make a slice for the values
    values := make([]sql.RawBytes, len(columns))

    // rows.Scan wants '[]interface{}' as an argument, so we must copy the
    // references into such a slice
    // See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
    scanArgs := make([]interface{}, len(values))

    for i := range values {
        //arr := strings.Split(columns[i],"`")
        //columns[i] = arr[0]
        //types[i],_ = strconv.Atoi(arr[1])
        scanArgs[i] = &values[i]
    }

    dict := make(map[string]interface{})
    dict["Columns"] = columns

    // Fetch rows
    for rows.Next() {
        // get RawBytes from data
        err = rows.Scan(scanArgs...)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
        //log.Println("values:",values)
        /*tmpbuf := make([]sql.RawBytes, len(values))
        copy(tmpbuf,values)
        data = append(data,tmpbuf)*/
        var tmpbuf []interface{}
        // Now do something with the data.
        // Here we just print each column as a string.
        var value string
        for _, col := range values {
            // Here we can check if the value is nil (NULL value)
            if col == nil {
                value = "NULL"
            } else {
                value = string(col)
            }
            //fmt.Println(columns[i], ": ", value)
            tmpbuf = append(tmpbuf,value)
        }
        //fmt.Println("-----------------------------------")
        data = append(data,tmpbuf)

    }
    if err = rows.Err(); err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    dict["Data"] = data
    return dict
}


func (conn *Connection)Delete(sqlStr string,args ...interface{}){
    if conn.db==nil {
        return;
    }
    if len(args) == 0 {

    }else{
      stmt,err := conn.db.Prepare(sqlStr);
      if err!=nil {
          fmt.Println(err.Error());
          return;
      }
      defer stmt.Close();
      if result,err := stmt.Exec(args);err==nil {
          if c,err :=  result.RowsAffected();err==nil{
              fmt.Println("remove count : ",c);
          }
      }
    }

}

func (conn *Connection)Close(){
    if conn.db!=nil {
        conn.db.Close();
    }
}
