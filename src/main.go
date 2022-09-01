package main

import (
	_ "Miyo_Assignment/global"
	v1 "Miyo_Assignment/router"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

func main() {
	router := v1.InitRouter()

	router.Run()
}
