package cmd

import (
	_ "oa-review/dao"
	router "oa-review/routers"
)

func Run() {
	r := router.NewRouter()
	r.Listen(":8080")
}
