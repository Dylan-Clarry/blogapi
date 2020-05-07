package main

import (
	"fmt"
	"github.com/dylan-clarry/blogapi/app"
	"github.com/dylan-clarry/blogapi/config"
)

func main() {
	fmt.Println("hello")

	config := config.CreateConfig()

	fmt.Println(config.DB.DBName)

	app := &app.App{}
	app.Initialize(config)
}