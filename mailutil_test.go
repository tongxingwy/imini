package imini

import(
  "testing"
  "log"
  )


func TestParseDataLayout(t *testing.T){
    log.Println(parseDataLayout("现在时间{{2006年 5月2日}}日{{%Y年 %m月%d日}}报"))
    log.Println(parseDataLayout("现在时间{{%Y年 %m月%d日}}日报"))
    log.Println(parseDataLayout("现在时间{{%y年 %m月%d日}}日报"))
    log.Println(parseDataLayout("现在时间%y年 %m月%d日日报"))
    log.Println(parseDataLayout("{{%Y年 %m月%d日}}"))
}
