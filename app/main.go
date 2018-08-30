package main

import "github.com/u22-2018/gae/router"

func main() {
	r := router.GetRouter()
	r.Run(":8000")
}
