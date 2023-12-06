package main

import (
	"auditlimit/api"
	"auditlimit/config"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

func main() {
	s := g.Server()
	s.SetPort(config.PORT)
	s.BindHandler("/", Index)
	s.BindHandler("/audit_limit", api.AuditLimit)
	s.Run()
}

func Index(r *ghttp.Request) {
	r.Response.Write("Hello Xyhelper,this is auditlimit")
}

func init() {
	// 判断 ./data/keywords.txt 文件是否存在
	if gfile.Exists("./data/keywords.txt") {
		// 读取 ./data/keywords.txt 文件内容
		keyWords := gfile.GetContents("./data/keywords.txt")
		// 将文件内容按照换行符分割为切片
		keyWordsSlice := gstr.Split(keyWords, "\n")
		// 去除切片中的空格及空行
		if len(keyWordsSlice) > 0 {
			for i := 0; i < len(keyWordsSlice); i++ {
				keyWordsSlice[i] = gstr.Trim(keyWordsSlice[i])
				if keyWordsSlice[i] == "" {
					keyWordsSlice = append(keyWordsSlice[:i], keyWordsSlice[i+1:]...)
					i--
				}
			}
		}
		config.ForbiddenWords = keyWordsSlice
	}
}
