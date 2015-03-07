package DriverManager

import(
  "time"
)

const (
  MYSQL_TYPE_BIT              =    16
  MYSQL_TYPE_BLOB             =    252
  MYSQL_TYPE_DATE             =    10
  MYSQL_TYPE_DATETIME         =    12
  MYSQL_TYPE_DECIMAL          =    0
  MYSQL_TYPE_DOUBLE           =    5
  MYSQL_TYPE_ENUM             =    247
  MYSQL_TYPE_FLOAT            =    4
  MYSQL_TYPE_GEOMETRY         =    255
  MYSQL_TYPE_INT24            =    9
  MYSQL_TYPE_LONG             =    3
  MYSQL_TYPE_LONGLONG         =    8
  MYSQL_TYPE_LONG_BLOB        =    251
  MYSQL_TYPE_MEDIUM_BLOB      =    250
  MYSQL_TYPE_NEWDATE          =    14
  MYSQL_TYPE_NEWDECIMAL       =    246
  MYSQL_TYPE_NULL             =    6
  MYSQL_TYPE_SET              =    248
  MYSQL_TYPE_SHORT            =    2
  MYSQL_TYPE_STRING           =    254
  MYSQL_TYPE_TIME             =    11
  MYSQL_TYPE_TIMESTAMP        =    7
  MYSQL_TYPE_TINY             =    1
  MYSQL_TYPE_TINY_BLOB        =    249
  MYSQL_TYPE_VARCHAR          =    15
  MYSQL_TYPE_VAR_STRING       =    253
  MYSQL_TYPE_YEAR             =    13
)

func getType(types int) interface{}{
  switch types{
    case MYSQL_TYPE_NULL,MYSQL_TYPE_BIT,MYSQL_TYPE_DECIMAL,MYSQL_TYPE_ENUM,MYSQL_TYPE_INT24,MYSQL_TYPE_LONG,MYSQL_TYPE_LONGLONG,MYSQL_TYPE_SHORT,MYSQL_TYPE_TINY:
      return int64(0)
    case MYSQL_TYPE_DATE,MYSQL_TYPE_DATETIME,MYSQL_TYPE_NEWDATE,MYSQL_TYPE_TIME,MYSQL_TYPE_TIMESTAMP,MYSQL_TYPE_YEAR:
      return time.Now()
    case MYSQL_TYPE_DOUBLE,MYSQL_TYPE_FLOAT:
      return float64(0.0)
    case MYSQL_TYPE_STRING,MYSQL_TYPE_VARCHAR,MYSQL_TYPE_VAR_STRING:
      return string("")
    default:
      return make([]byte,1,1)
    }
}
