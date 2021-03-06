package main

import (
    "gitee.com/johng/gf/g/net/ghttp"
)

func main() {
    s := ghttp.GetServer()
    s.BindHandler("/", func(r *ghttp.Request){
        r.Response.Writeln("您可以同时通过HTTP和HTTPS方式看到该内容！")
    })
    s.EnableHTTPS("/home/john/temp/server.crt", "/home/john/temp/server.key")
    s.SetHTTPSPort(443)
    s.SetPort(80)
    s.Run()
}