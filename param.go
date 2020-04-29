package main

import "flag"

var (
	Uri		string
	Date	string
)

func init()  {
	flag.StringVar(&Uri, "uri", "", "请求地址")
	flag.StringVar(&Date, "date", "", "日期")
}
