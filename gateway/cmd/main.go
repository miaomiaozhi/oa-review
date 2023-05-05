package cmd

import "oa-review/gateway/router"

func Run() {
	r := router.NewRouter()
	r.Listen(":8080")
}
