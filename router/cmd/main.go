package cmd

import (
	_ "oa-review/dao"
	"oa-review/router"
)

func Run() {
	r := router.NewRouter()
	r.Listen(":8080")
}
