package DateUtil

import(
  "log"
  "testing"
  )

func TestParseDate(t *testing.T){
  log.Println(FormatNow("yyyy-mm-dd")) //not support
  log.Println(FormatNow("%Y年 %m月%d日")) //like c style strftime
}
