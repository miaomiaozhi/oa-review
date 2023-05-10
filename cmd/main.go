package main

import (
	_ "oa-review/conf"
	dao "oa-review/dao"
	gateway "oa-review/gateway/cmd"
)

func main() {
	if err := dao.InitDataBase(); err != nil {
		return
	}
	gateway.Run()
}
